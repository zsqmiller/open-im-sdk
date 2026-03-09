package initialization

import (
	"github.com/zsqmiller/open-im-sdk/v3/integration_test/internal/manager"
)

func GenUserIDs() {
	manager.NewUserManager(manager.NewMetaManager()).GenAllUserIDs()
}
