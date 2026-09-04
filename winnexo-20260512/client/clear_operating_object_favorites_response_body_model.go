// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearOperatingObjectFavoritesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ClearOperatingObjectFavoritesResponseBody
	GetCode() *string
	SetGraphName(v string) *ClearOperatingObjectFavoritesResponseBody
	GetGraphName() *string
	SetMessage(v string) *ClearOperatingObjectFavoritesResponseBody
	GetMessage() *string
	SetObjectType(v string) *ClearOperatingObjectFavoritesResponseBody
	GetObjectType() *string
	SetOperatingObjectName(v string) *ClearOperatingObjectFavoritesResponseBody
	GetOperatingObjectName() *string
	SetRemainingCount(v int64) *ClearOperatingObjectFavoritesResponseBody
	GetRemainingCount() *int64
	SetRemovedCount(v int64) *ClearOperatingObjectFavoritesResponseBody
	GetRemovedCount() *int64
	SetRequestId(v string) *ClearOperatingObjectFavoritesResponseBody
	GetRequestId() *string
	SetVerified(v bool) *ClearOperatingObjectFavoritesResponseBody
	GetVerified() *bool
}

type ClearOperatingObjectFavoritesResponseBody struct {
	// The status code. SUCCESS indicates success. In case of failure, the corresponding error type is returned, such as ERR_BAD_REQUEST, ERR_VALIDATION_FAILED, or ERR_INTERNAL_SERVER_ERROR.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The graph name. You can call listGraphs to obtain the value.
	//
	// example:
	//
	// crm
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The status code description.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The object type, such as customer. This parameter has a value when type is set to mention.
	//
	// example:
	//
	// contract
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The digital employee name (operating object name, optional).
	//
	// example:
	//
	// customer_assistant
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The number of remaining followed objects within the specified scope.
	//
	// example:
	//
	// 0
	RemainingCount *int64 `json:"remainingCount,omitempty" xml:"remainingCount,omitempty"`
	// The number of physical follow records that were actually deleted.
	//
	// example:
	//
	// 2
	RemovedCount *int64 `json:"removedCount,omitempty" xml:"removedCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the remaining record count has been verified as zero within the same transaction.
	//
	// example:
	//
	// true
	Verified *bool `json:"verified,omitempty" xml:"verified,omitempty"`
}

func (s ClearOperatingObjectFavoritesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearOperatingObjectFavoritesResponseBody) GoString() string {
	return s.String()
}

func (s *ClearOperatingObjectFavoritesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClearOperatingObjectFavoritesResponseBody) GetGraphName() *string {
	return s.GraphName
}

func (s *ClearOperatingObjectFavoritesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClearOperatingObjectFavoritesResponseBody) GetObjectType() *string {
	return s.ObjectType
}

func (s *ClearOperatingObjectFavoritesResponseBody) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ClearOperatingObjectFavoritesResponseBody) GetRemainingCount() *int64 {
	return s.RemainingCount
}

func (s *ClearOperatingObjectFavoritesResponseBody) GetRemovedCount() *int64 {
	return s.RemovedCount
}

func (s *ClearOperatingObjectFavoritesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClearOperatingObjectFavoritesResponseBody) GetVerified() *bool {
	return s.Verified
}

func (s *ClearOperatingObjectFavoritesResponseBody) SetCode(v string) *ClearOperatingObjectFavoritesResponseBody {
	s.Code = &v
	return s
}

func (s *ClearOperatingObjectFavoritesResponseBody) SetGraphName(v string) *ClearOperatingObjectFavoritesResponseBody {
	s.GraphName = &v
	return s
}

func (s *ClearOperatingObjectFavoritesResponseBody) SetMessage(v string) *ClearOperatingObjectFavoritesResponseBody {
	s.Message = &v
	return s
}

func (s *ClearOperatingObjectFavoritesResponseBody) SetObjectType(v string) *ClearOperatingObjectFavoritesResponseBody {
	s.ObjectType = &v
	return s
}

func (s *ClearOperatingObjectFavoritesResponseBody) SetOperatingObjectName(v string) *ClearOperatingObjectFavoritesResponseBody {
	s.OperatingObjectName = &v
	return s
}

func (s *ClearOperatingObjectFavoritesResponseBody) SetRemainingCount(v int64) *ClearOperatingObjectFavoritesResponseBody {
	s.RemainingCount = &v
	return s
}

func (s *ClearOperatingObjectFavoritesResponseBody) SetRemovedCount(v int64) *ClearOperatingObjectFavoritesResponseBody {
	s.RemovedCount = &v
	return s
}

func (s *ClearOperatingObjectFavoritesResponseBody) SetRequestId(v string) *ClearOperatingObjectFavoritesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearOperatingObjectFavoritesResponseBody) SetVerified(v bool) *ClearOperatingObjectFavoritesResponseBody {
	s.Verified = &v
	return s
}

func (s *ClearOperatingObjectFavoritesResponseBody) Validate() error {
	return dara.Validate(s)
}
