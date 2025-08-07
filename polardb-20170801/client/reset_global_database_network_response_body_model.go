// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetGlobalDatabaseNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetGlobalDatabaseNetworkResponseBody
	GetRequestId() *string
}

type ResetGlobalDatabaseNetworkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 67F2E75F-AE67-4FB2-821F-A81237EACD15
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetGlobalDatabaseNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetGlobalDatabaseNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *ResetGlobalDatabaseNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetGlobalDatabaseNetworkResponseBody) SetRequestId(v string) *ResetGlobalDatabaseNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetGlobalDatabaseNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
