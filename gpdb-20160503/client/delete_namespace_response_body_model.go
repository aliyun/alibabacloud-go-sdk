// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteNamespaceResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteNamespaceResponseBody
	GetStatus() *string
}

type DeleteNamespaceResponseBody struct {
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

func (s DeleteNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNamespaceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteNamespaceResponseBody) SetMessage(v string) *DeleteNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetRequestId(v string) *DeleteNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetStatus(v string) *DeleteNamespaceResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
