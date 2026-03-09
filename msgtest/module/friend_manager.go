package module

import (
	"github.com/zsqmiller/open-im-sdk/v3/pkg/api"
	"github.com/openimsdk/protocol/relation"
)

type TestFriendManager struct {
	*MetaManager
}

func (t *TestFriendManager) ImportFriends(ownerUserID string, friendUserIDs []string) error {
	req := &relation.ImportFriendReq{
		OwnerUserID:   ownerUserID,
		FriendUserIDs: friendUserIDs,
	}
	return t.postWithCtx(api.ImportFriendList.Route(), &req, nil)
}
