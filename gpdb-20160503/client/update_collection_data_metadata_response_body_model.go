// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCollectionDataMetadataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppliedRows(v int64) *UpdateCollectionDataMetadataResponseBody
	GetAppliedRows() *int64
	SetMessage(v string) *UpdateCollectionDataMetadataResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCollectionDataMetadataResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateCollectionDataMetadataResponseBody
	GetStatus() *string
}

type UpdateCollectionDataMetadataResponseBody struct {
	// Number of effective entries.
	//
	// example:
	//
	// 10
	AppliedRows *int64 `json:"AppliedRows,omitempty" xml:"AppliedRows,omitempty"`
	// Detailed information when the request fails.
	//
	// example:
	//
	// failed to connect database, detailMsg: getConnection fail::SQL State: 28P01, Error Code: 0, Error Message: FATAL: password
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Status, with the following values:
	//
	// - **success**: Success.
	//
	// - **fail**: Failure.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateCollectionDataMetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCollectionDataMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCollectionDataMetadataResponseBody) GetAppliedRows() *int64 {
	return s.AppliedRows
}

func (s *UpdateCollectionDataMetadataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCollectionDataMetadataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCollectionDataMetadataResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateCollectionDataMetadataResponseBody) SetAppliedRows(v int64) *UpdateCollectionDataMetadataResponseBody {
	s.AppliedRows = &v
	return s
}

func (s *UpdateCollectionDataMetadataResponseBody) SetMessage(v string) *UpdateCollectionDataMetadataResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCollectionDataMetadataResponseBody) SetRequestId(v string) *UpdateCollectionDataMetadataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCollectionDataMetadataResponseBody) SetStatus(v string) *UpdateCollectionDataMetadataResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateCollectionDataMetadataResponseBody) Validate() error {
	return dara.Validate(s)
}
