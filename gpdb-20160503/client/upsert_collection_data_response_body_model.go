// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertCollectionDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpsertCollectionDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpsertCollectionDataResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpsertCollectionDataResponseBody
	GetStatus() *string
}

type UpsertCollectionDataResponseBody struct {
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

func (s UpsertCollectionDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpsertCollectionDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpsertCollectionDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpsertCollectionDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpsertCollectionDataResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpsertCollectionDataResponseBody) SetMessage(v string) *UpsertCollectionDataResponseBody {
	s.Message = &v
	return s
}

func (s *UpsertCollectionDataResponseBody) SetRequestId(v string) *UpsertCollectionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpsertCollectionDataResponseBody) SetStatus(v string) *UpsertCollectionDataResponseBody {
	s.Status = &v
	return s
}

func (s *UpsertCollectionDataResponseBody) Validate() error {
	return dara.Validate(s)
}
