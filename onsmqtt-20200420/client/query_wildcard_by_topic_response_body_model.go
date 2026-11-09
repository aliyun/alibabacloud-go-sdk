// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWildcardByTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryWildcardByTopicResponseBody
	GetCode() *int32
	SetData(v *QueryWildcardByTopicResponseBodyData) *QueryWildcardByTopicResponseBody
	GetData() *QueryWildcardByTopicResponseBodyData
	SetRequestId(v string) *QueryWildcardByTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryWildcardByTopicResponseBody
	GetSuccess() *bool
}

type QueryWildcardByTopicResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dataset of parent topic information.
	Data *QueryWildcardByTopicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 8CC04203-679B-4DED-89D9-E7C2E979****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryWildcardByTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryWildcardByTopicResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWildcardByTopicResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryWildcardByTopicResponseBody) GetData() *QueryWildcardByTopicResponseBodyData {
	return s.Data
}

func (s *QueryWildcardByTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryWildcardByTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryWildcardByTopicResponseBody) SetCode(v int32) *QueryWildcardByTopicResponseBody {
	s.Code = &v
	return s
}

func (s *QueryWildcardByTopicResponseBody) SetData(v *QueryWildcardByTopicResponseBodyData) *QueryWildcardByTopicResponseBody {
	s.Data = v
	return s
}

func (s *QueryWildcardByTopicResponseBody) SetRequestId(v string) *QueryWildcardByTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWildcardByTopicResponseBody) SetSuccess(v bool) *QueryWildcardByTopicResponseBody {
	s.Success = &v
	return s
}

func (s *QueryWildcardByTopicResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryWildcardByTopicResponseBodyData struct {
	// The maximum number of results returned.
	//
	// example:
	//
	// 20
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The total number of query results.
	//
	// example:
	//
	// 1
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
	// The wildcard list.
	WildcardList []*string `json:"WildcardList,omitempty" xml:"WildcardList,omitempty" type:"Repeated"`
}

func (s QueryWildcardByTopicResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryWildcardByTopicResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryWildcardByTopicResponseBodyData) GetLimit() *string {
	return s.Limit
}

func (s *QueryWildcardByTopicResponseBodyData) GetTotal() *string {
	return s.Total
}

func (s *QueryWildcardByTopicResponseBodyData) GetWildcardList() []*string {
	return s.WildcardList
}

func (s *QueryWildcardByTopicResponseBodyData) SetLimit(v string) *QueryWildcardByTopicResponseBodyData {
	s.Limit = &v
	return s
}

func (s *QueryWildcardByTopicResponseBodyData) SetTotal(v string) *QueryWildcardByTopicResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryWildcardByTopicResponseBodyData) SetWildcardList(v []*string) *QueryWildcardByTopicResponseBodyData {
	s.WildcardList = v
	return s
}

func (s *QueryWildcardByTopicResponseBodyData) Validate() error {
	return dara.Validate(s)
}
