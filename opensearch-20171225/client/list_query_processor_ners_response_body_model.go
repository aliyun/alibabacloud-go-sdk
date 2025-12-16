// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueryProcessorNersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListQueryProcessorNersResponseBody
	GetRequestId() *string
	SetResult(v []*ListQueryProcessorNersResponseBodyResult) *ListQueryProcessorNersResponseBody
	GetResult() []*ListQueryProcessorNersResponseBodyResult
}

type ListQueryProcessorNersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The priority settings of entity types.
	//
	// For more information, see [Priority settings of entity types](https://help.aliyun.com/document_detail/173616.html).
	Result []*ListQueryProcessorNersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListQueryProcessorNersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQueryProcessorNersResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorNersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQueryProcessorNersResponseBody) GetResult() []*ListQueryProcessorNersResponseBodyResult {
	return s.Result
}

func (s *ListQueryProcessorNersResponseBody) SetRequestId(v string) *ListQueryProcessorNersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueryProcessorNersResponseBody) SetResult(v []*ListQueryProcessorNersResponseBodyResult) *ListQueryProcessorNersResponseBody {
	s.Result = v
	return s
}

func (s *ListQueryProcessorNersResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQueryProcessorNersResponseBodyResult struct {
	// The name of the entity type.
	//
	// example:
	//
	// brand
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// The priority of an entity type among entity types that have the same priority level. A smaller value indicates a higher priority. Default value: 0.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// The priority level of the entity type. Valid values:
	//
	// 	- HIGH
	//
	// 	- MIDDLE
	//
	// 	- LOW
	//
	// example:
	//
	// HIGH
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// The internal name of the entity type.
	//
	// example:
	//
	// brand
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListQueryProcessorNersResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListQueryProcessorNersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorNersResponseBodyResult) GetLabel() *string {
	return s.Label
}

func (s *ListQueryProcessorNersResponseBodyResult) GetOrder() *int32 {
	return s.Order
}

func (s *ListQueryProcessorNersResponseBodyResult) GetPriority() *string {
	return s.Priority
}

func (s *ListQueryProcessorNersResponseBodyResult) GetTag() *string {
	return s.Tag
}

func (s *ListQueryProcessorNersResponseBodyResult) SetLabel(v string) *ListQueryProcessorNersResponseBodyResult {
	s.Label = &v
	return s
}

func (s *ListQueryProcessorNersResponseBodyResult) SetOrder(v int32) *ListQueryProcessorNersResponseBodyResult {
	s.Order = &v
	return s
}

func (s *ListQueryProcessorNersResponseBodyResult) SetPriority(v string) *ListQueryProcessorNersResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *ListQueryProcessorNersResponseBodyResult) SetTag(v string) *ListQueryProcessorNersResponseBodyResult {
	s.Tag = &v
	return s
}

func (s *ListQueryProcessorNersResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
