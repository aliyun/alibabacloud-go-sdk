// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCollectionDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppliedRows(v int64) *DeleteCollectionDataResponseBody
	GetAppliedRows() *int64
	SetMessage(v string) *DeleteCollectionDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCollectionDataResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteCollectionDataResponseBody
	GetStatus() *string
}

type DeleteCollectionDataResponseBody struct {
	// The number of rows that are affected by the request.
	//
	// example:
	//
	// 10
	AppliedRows *int64 `json:"AppliedRows,omitempty" xml:"AppliedRows,omitempty"`
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

func (s DeleteCollectionDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCollectionDataResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCollectionDataResponseBody) GetAppliedRows() *int64 {
	return s.AppliedRows
}

func (s *DeleteCollectionDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCollectionDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCollectionDataResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteCollectionDataResponseBody) SetAppliedRows(v int64) *DeleteCollectionDataResponseBody {
	s.AppliedRows = &v
	return s
}

func (s *DeleteCollectionDataResponseBody) SetMessage(v string) *DeleteCollectionDataResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCollectionDataResponseBody) SetRequestId(v string) *DeleteCollectionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCollectionDataResponseBody) SetStatus(v string) *DeleteCollectionDataResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteCollectionDataResponseBody) Validate() error {
	return dara.Validate(s)
}
