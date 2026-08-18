// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterSearchClientTreeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterSearchClientTreeResponseBodyData) *ModelRouterSearchClientTreeResponseBody
	GetData() *ModelRouterSearchClientTreeResponseBodyData
	SetErrCode(v string) *ModelRouterSearchClientTreeResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterSearchClientTreeResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterSearchClientTreeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterSearchClientTreeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterSearchClientTreeResponseBody
	GetSuccess() *bool
}

type ModelRouterSearchClientTreeResponseBody struct {
	// The response data object.
	//
	// example:
	//
	// {}
	Data *ModelRouterSearchClientTreeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The fault message code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterSearchClientTreeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterSearchClientTreeResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterSearchClientTreeResponseBody) GetData() *ModelRouterSearchClientTreeResponseBodyData {
	return s.Data
}

func (s *ModelRouterSearchClientTreeResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterSearchClientTreeResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterSearchClientTreeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterSearchClientTreeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterSearchClientTreeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterSearchClientTreeResponseBody) SetData(v *ModelRouterSearchClientTreeResponseBodyData) *ModelRouterSearchClientTreeResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterSearchClientTreeResponseBody) SetErrCode(v string) *ModelRouterSearchClientTreeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterSearchClientTreeResponseBody) SetErrMessage(v string) *ModelRouterSearchClientTreeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterSearchClientTreeResponseBody) SetHttpStatusCode(v int32) *ModelRouterSearchClientTreeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterSearchClientTreeResponseBody) SetRequestId(v string) *ModelRouterSearchClientTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterSearchClientTreeResponseBody) SetSuccess(v bool) *ModelRouterSearchClientTreeResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterSearchClientTreeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterSearchClientTreeResponseBodyData struct {
	// The collection of matched department IDs.
	//
	// example:
	//
	// []
	MatchedDeptIds []*int64 `json:"matchedDeptIds,omitempty" xml:"matchedDeptIds,omitempty" type:"Repeated"`
}

func (s ModelRouterSearchClientTreeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterSearchClientTreeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterSearchClientTreeResponseBodyData) GetMatchedDeptIds() []*int64 {
	return s.MatchedDeptIds
}

func (s *ModelRouterSearchClientTreeResponseBodyData) SetMatchedDeptIds(v []*int64) *ModelRouterSearchClientTreeResponseBodyData {
	s.MatchedDeptIds = v
	return s
}

func (s *ModelRouterSearchClientTreeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
