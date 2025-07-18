// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *GrantCollectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GrantCollectionResponseBody
	GetRequestId() *string
	SetStatus(v string) *GrantCollectionResponseBody
	GetStatus() *string
}

type GrantCollectionResponseBody struct {
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

func (s GrantCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *GrantCollectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GrantCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantCollectionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GrantCollectionResponseBody) SetMessage(v string) *GrantCollectionResponseBody {
	s.Message = &v
	return s
}

func (s *GrantCollectionResponseBody) SetRequestId(v string) *GrantCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantCollectionResponseBody) SetStatus(v string) *GrantCollectionResponseBody {
	s.Status = &v
	return s
}

func (s *GrantCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
