// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDevelopServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DevelopServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DevelopServiceResponseBody
	GetRequestId() *string
}

type DevelopServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DevelopServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DevelopServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DevelopServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DevelopServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DevelopServiceResponseBody) SetMessage(v string) *DevelopServiceResponseBody {
	s.Message = &v
	return s
}

func (s *DevelopServiceResponseBody) SetRequestId(v string) *DevelopServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DevelopServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
