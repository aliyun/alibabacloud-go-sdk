// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAppKeyResponseBody
	GetRequestId() *string
	SetData(v bool) *CreateAppKeyResponseBody
	GetData() *bool
}

type CreateAppKeyResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Data object
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateAppKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppKeyResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateAppKeyResponseBody) SetRequestId(v string) *CreateAppKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppKeyResponseBody) SetData(v bool) *CreateAppKeyResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAppKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
