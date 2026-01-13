// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecallManagementConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRecallManagementConfigResponseBody
	GetRequestId() *string
}

type CreateRecallManagementConfigResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRecallManagementConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRecallManagementConfigResponseBody) SetRequestId(v string) *CreateRecallManagementConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecallManagementConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
