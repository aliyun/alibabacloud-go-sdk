// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *RenameApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *RenameApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RenameApiKeyResponseBody
	GetSuccess() *bool
}

type RenameApiKeyResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenameApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenameApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *RenameApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RenameApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenameApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RenameApiKeyResponseBody) SetMessage(v string) *RenameApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *RenameApiKeyResponseBody) SetRequestId(v string) *RenameApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameApiKeyResponseBody) SetSuccess(v bool) *RenameApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *RenameApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
