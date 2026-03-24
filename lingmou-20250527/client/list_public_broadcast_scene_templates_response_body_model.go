// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicBroadcastSceneTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPublicBroadcastSceneTemplatesResponseBody
	GetCode() *string
	SetData(v []*BroadcastSceneTemplate) *ListPublicBroadcastSceneTemplatesResponseBody
	GetData() []*BroadcastSceneTemplate
	SetHttpStatusCode(v int32) *ListPublicBroadcastSceneTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListPublicBroadcastSceneTemplatesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListPublicBroadcastSceneTemplatesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListPublicBroadcastSceneTemplatesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPublicBroadcastSceneTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPublicBroadcastSceneTemplatesResponseBody
	GetSuccess() *bool
}

type ListPublicBroadcastSceneTemplatesResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                   `json:"code,omitempty" xml:"code,omitempty"`
	Data []*BroadcastSceneTemplate `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	MaxResults     *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// SUCCESS
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 7239F9E5-A4DB-55BA-B701-0CE47DBDB0A8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListPublicBroadcastSceneTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPublicBroadcastSceneTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicBroadcastSceneTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPublicBroadcastSceneTemplatesResponseBody) GetData() []*BroadcastSceneTemplate {
	return s.Data
}

func (s *ListPublicBroadcastSceneTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPublicBroadcastSceneTemplatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPublicBroadcastSceneTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPublicBroadcastSceneTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPublicBroadcastSceneTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPublicBroadcastSceneTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPublicBroadcastSceneTemplatesResponseBody) SetCode(v string) *ListPublicBroadcastSceneTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPublicBroadcastSceneTemplatesResponseBody) SetData(v []*BroadcastSceneTemplate) *ListPublicBroadcastSceneTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListPublicBroadcastSceneTemplatesResponseBody) SetHttpStatusCode(v int32) *ListPublicBroadcastSceneTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPublicBroadcastSceneTemplatesResponseBody) SetMaxResults(v int32) *ListPublicBroadcastSceneTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPublicBroadcastSceneTemplatesResponseBody) SetMessage(v string) *ListPublicBroadcastSceneTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPublicBroadcastSceneTemplatesResponseBody) SetNextToken(v string) *ListPublicBroadcastSceneTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicBroadcastSceneTemplatesResponseBody) SetRequestId(v string) *ListPublicBroadcastSceneTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicBroadcastSceneTemplatesResponseBody) SetSuccess(v bool) *ListPublicBroadcastSceneTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListPublicBroadcastSceneTemplatesResponseBody) Validate() error {
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
