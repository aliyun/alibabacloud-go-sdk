// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateNamespaceResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateNamespaceResponseBody
	GetStatus() *string
}

type CreateNamespaceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNamespaceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateNamespaceResponseBody) SetMessage(v string) *CreateNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetRequestId(v string) *CreateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetStatus(v string) *CreateNamespaceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
