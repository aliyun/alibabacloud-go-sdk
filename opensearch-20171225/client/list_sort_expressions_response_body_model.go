// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSortExpressionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSortExpressionsResponseBody
	GetRequestId() *string
	SetResult(v []*ListSortExpressionsResponseBodyResult) *ListSortExpressionsResponseBody
	GetResult() []*ListSortExpressionsResponseBodyResult
}

type ListSortExpressionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the rough sort or fine sort expressions that are returned.
	//
	// For more information, see [FirstRank](https://help.aliyun.com/document_detail/170007.html) and [SecondRank](https://help.aliyun.com/document_detail/170008.html).
	Result []*ListSortExpressionsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListSortExpressionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSortExpressionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSortExpressionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSortExpressionsResponseBody) GetResult() []*ListSortExpressionsResponseBodyResult {
	return s.Result
}

func (s *ListSortExpressionsResponseBody) SetRequestId(v string) *ListSortExpressionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSortExpressionsResponseBody) SetResult(v []*ListSortExpressionsResponseBodyResult) *ListSortExpressionsResponseBody {
	s.Result = v
	return s
}

func (s *ListSortExpressionsResponseBody) Validate() error {
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

type ListSortExpressionsResponseBodyResult struct {
	// Indicates whether the expression is the default one.
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The timestamp when the sort expression was created.
	//
	// example:
	//
	// 1655793690
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The description of the sort expression.
	//
	// example:
	//
	// ""
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the sort expression.
	//
	// example:
	//
	// default
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The timestamp when the sort expression was updated.
	//
	// example:
	//
	// 1655793690
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListSortExpressionsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListSortExpressionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSortExpressionsResponseBodyResult) GetActive() *bool {
	return s.Active
}

func (s *ListSortExpressionsResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *ListSortExpressionsResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListSortExpressionsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListSortExpressionsResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *ListSortExpressionsResponseBodyResult) SetActive(v bool) *ListSortExpressionsResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ListSortExpressionsResponseBodyResult) SetCreated(v int32) *ListSortExpressionsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListSortExpressionsResponseBodyResult) SetDescription(v string) *ListSortExpressionsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListSortExpressionsResponseBodyResult) SetName(v string) *ListSortExpressionsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListSortExpressionsResponseBodyResult) SetUpdated(v int32) *ListSortExpressionsResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListSortExpressionsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
