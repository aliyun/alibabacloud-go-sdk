// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchOverGlobalDatabaseNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchOverGlobalDatabaseNetworkResponseBody
	GetRequestId() *string
}

type SwitchOverGlobalDatabaseNetworkResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 67F2E75F-AE67-4FB2-821F-A81237EACD15
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchOverGlobalDatabaseNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchOverGlobalDatabaseNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchOverGlobalDatabaseNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchOverGlobalDatabaseNetworkResponseBody) SetRequestId(v string) *SwitchOverGlobalDatabaseNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchOverGlobalDatabaseNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
