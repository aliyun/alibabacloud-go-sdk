// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeRecallManagementServiceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeRecallManagementServiceVersionResponseBody
	GetRequestId() *string
}

type ChangeRecallManagementServiceVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeRecallManagementServiceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeRecallManagementServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeRecallManagementServiceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeRecallManagementServiceVersionResponseBody) SetRequestId(v string) *ChangeRecallManagementServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeRecallManagementServiceVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
