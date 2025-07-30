// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizationConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCustomizationConfigResponseBody
	GetCode() *string
	SetData(v string) *DeleteCustomizationConfigResponseBody
	GetData() *string
	SetMessage(v string) *DeleteCustomizationConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCustomizationConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCustomizationConfigResponseBody
	GetSuccess() *bool
}

type DeleteCustomizationConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 252
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomizationConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomizationConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCustomizationConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteCustomizationConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomizationConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomizationConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomizationConfigResponseBody) SetCode(v string) *DeleteCustomizationConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomizationConfigResponseBody) SetData(v string) *DeleteCustomizationConfigResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCustomizationConfigResponseBody) SetMessage(v string) *DeleteCustomizationConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomizationConfigResponseBody) SetRequestId(v string) *DeleteCustomizationConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomizationConfigResponseBody) SetSuccess(v bool) *DeleteCustomizationConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomizationConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
