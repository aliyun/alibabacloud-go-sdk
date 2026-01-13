// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineRecallManagementServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OnlineRecallManagementServiceResponseBody
	GetRequestId() *string
}

type OnlineRecallManagementServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnlineRecallManagementServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnlineRecallManagementServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OnlineRecallManagementServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnlineRecallManagementServiceResponseBody) SetRequestId(v string) *OnlineRecallManagementServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnlineRecallManagementServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
