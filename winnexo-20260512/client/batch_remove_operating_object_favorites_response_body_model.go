// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRemoveOperatingObjectFavoritesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchRemoveOperatingObjectFavoritesResponseBody
	GetCode() *string
	SetGraphName(v string) *BatchRemoveOperatingObjectFavoritesResponseBody
	GetGraphName() *string
	SetMessage(v string) *BatchRemoveOperatingObjectFavoritesResponseBody
	GetMessage() *string
	SetObjectType(v string) *BatchRemoveOperatingObjectFavoritesResponseBody
	GetObjectType() *string
	SetOperatingObjectName(v string) *BatchRemoveOperatingObjectFavoritesResponseBody
	GetOperatingObjectName() *string
	SetRemainingCount(v int64) *BatchRemoveOperatingObjectFavoritesResponseBody
	GetRemainingCount() *int64
	SetRemovedCount(v int64) *BatchRemoveOperatingObjectFavoritesResponseBody
	GetRemovedCount() *int64
	SetRequestId(v string) *BatchRemoveOperatingObjectFavoritesResponseBody
	GetRequestId() *string
	SetRequestedCount(v int64) *BatchRemoveOperatingObjectFavoritesResponseBody
	GetRequestedCount() *int64
	SetResults(v []*BatchRemoveOperatingObjectFavoritesResponseBodyResults) *BatchRemoveOperatingObjectFavoritesResponseBody
	GetResults() []*BatchRemoveOperatingObjectFavoritesResponseBodyResults
}

type BatchRemoveOperatingObjectFavoritesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The graph name.
	//
	// example:
	//
	// string_value
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The object type, such as customer. This parameter has a value when type is set to mention.
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The digital employee name (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The number of remaining favorited objects within the specified scope.
	//
	// example:
	//
	// 0
	RemainingCount *int64 `json:"remainingCount,omitempty" xml:"remainingCount,omitempty"`
	// The number of physical favorite records that are actually deleted.
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
	// **The number of requested members before deduplication.**
	//
	// example:
	//
	// 2
	RequestedCount *int64 `json:"requestedCount,omitempty" xml:"requestedCount,omitempty"`
	// The relationships between internal and external DingTalk users that failed to be created.
	Results []*BatchRemoveOperatingObjectFavoritesResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s BatchRemoveOperatingObjectFavoritesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchRemoveOperatingObjectFavoritesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) GetGraphName() *string {
	return s.GraphName
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) GetObjectType() *string {
	return s.ObjectType
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) GetRemainingCount() *int64 {
	return s.RemainingCount
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) GetRemovedCount() *int64 {
	return s.RemovedCount
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) GetRequestedCount() *int64 {
	return s.RequestedCount
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) GetResults() []*BatchRemoveOperatingObjectFavoritesResponseBodyResults {
	return s.Results
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) SetCode(v string) *BatchRemoveOperatingObjectFavoritesResponseBody {
	s.Code = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) SetGraphName(v string) *BatchRemoveOperatingObjectFavoritesResponseBody {
	s.GraphName = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) SetMessage(v string) *BatchRemoveOperatingObjectFavoritesResponseBody {
	s.Message = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) SetObjectType(v string) *BatchRemoveOperatingObjectFavoritesResponseBody {
	s.ObjectType = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) SetOperatingObjectName(v string) *BatchRemoveOperatingObjectFavoritesResponseBody {
	s.OperatingObjectName = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) SetRemainingCount(v int64) *BatchRemoveOperatingObjectFavoritesResponseBody {
	s.RemainingCount = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) SetRemovedCount(v int64) *BatchRemoveOperatingObjectFavoritesResponseBody {
	s.RemovedCount = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) SetRequestId(v string) *BatchRemoveOperatingObjectFavoritesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) SetRequestedCount(v int64) *BatchRemoveOperatingObjectFavoritesResponseBody {
	s.RequestedCount = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) SetResults(v []*BatchRemoveOperatingObjectFavoritesResponseBodyResults) *BatchRemoveOperatingObjectFavoritesResponseBody {
	s.Results = v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchRemoveOperatingObjectFavoritesResponseBodyResults struct {
	// Indicates whether the object is favorited after the operation.
	//
	// example:
	//
	// true
	IsFavorited *bool `json:"isFavorited,omitempty" xml:"isFavorited,omitempty"`
	// The aligned object ID: target ID or KR ID.
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// Indicates whether the request has been processed.
	//
	// example:
	//
	// true
	Processed *bool `json:"processed,omitempty" xml:"processed,omitempty"`
}

func (s BatchRemoveOperatingObjectFavoritesResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchRemoveOperatingObjectFavoritesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBodyResults) GetIsFavorited() *bool {
	return s.IsFavorited
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBodyResults) GetObjectId() *string {
	return s.ObjectId
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBodyResults) GetProcessed() *bool {
	return s.Processed
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBodyResults) SetIsFavorited(v bool) *BatchRemoveOperatingObjectFavoritesResponseBodyResults {
	s.IsFavorited = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBodyResults) SetObjectId(v string) *BatchRemoveOperatingObjectFavoritesResponseBodyResults {
	s.ObjectId = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBodyResults) SetProcessed(v bool) *BatchRemoveOperatingObjectFavoritesResponseBodyResults {
	s.Processed = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
