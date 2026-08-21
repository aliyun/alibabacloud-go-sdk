// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstanceInfoResponseBody
	GetCode() *string
	SetData(v []*ListInstanceInfoResponseBodyData) *ListInstanceInfoResponseBody
	GetData() []*ListInstanceInfoResponseBodyData
	SetMaxResults(v int32) *ListInstanceInfoResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListInstanceInfoResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListInstanceInfoResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListInstanceInfoResponseBody
	GetRequestId() *string
}

type ListInstanceInfoResponseBody struct {
	// example:
	//
	// Success
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListInstanceInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// U+w1wv2R4ZWR5oZLXD0+Dp4dD+2BRJj42DLT6GrZysw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListInstanceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstanceInfoResponseBody) GetData() []*ListInstanceInfoResponseBodyData {
	return s.Data
}

func (s *ListInstanceInfoResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstanceInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstanceInfoResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstanceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceInfoResponseBody) SetCode(v string) *ListInstanceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceInfoResponseBody) SetData(v []*ListInstanceInfoResponseBodyData) *ListInstanceInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceInfoResponseBody) SetMaxResults(v int32) *ListInstanceInfoResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstanceInfoResponseBody) SetMessage(v string) *ListInstanceInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceInfoResponseBody) SetNextToken(v string) *ListInstanceInfoResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstanceInfoResponseBody) SetRequestId(v string) *ListInstanceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceInfoResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceInfoResponseBodyData struct {
	// example:
	//
	// sysom
	InfoKey *string `json:"infoKey,omitempty" xml:"infoKey,omitempty"`
	// example:
	//
	// instance_tag
	InfoType *string `json:"infoType,omitempty" xml:"infoType,omitempty"`
	// example:
	//
	// diagnosis
	InfoValue *string `json:"infoValue,omitempty" xml:"infoValue,omitempty"`
}

func (s ListInstanceInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceInfoResponseBodyData) GetInfoKey() *string {
	return s.InfoKey
}

func (s *ListInstanceInfoResponseBodyData) GetInfoType() *string {
	return s.InfoType
}

func (s *ListInstanceInfoResponseBodyData) GetInfoValue() *string {
	return s.InfoValue
}

func (s *ListInstanceInfoResponseBodyData) SetInfoKey(v string) *ListInstanceInfoResponseBodyData {
	s.InfoKey = &v
	return s
}

func (s *ListInstanceInfoResponseBodyData) SetInfoType(v string) *ListInstanceInfoResponseBodyData {
	s.InfoType = &v
	return s
}

func (s *ListInstanceInfoResponseBodyData) SetInfoValue(v string) *ListInstanceInfoResponseBodyData {
	s.InfoValue = &v
	return s
}

func (s *ListInstanceInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
