// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySearchIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySearchIndexResponseBody
	GetCode() *string
	SetIndexStatus(v string) *QuerySearchIndexResponseBody
	GetIndexStatus() *string
	SetIndexType(v string) *QuerySearchIndexResponseBody
	GetIndexType() *string
	SetMediaTotal(v string) *QuerySearchIndexResponseBody
	GetMediaTotal() *string
	SetRequestId(v string) *QuerySearchIndexResponseBody
	GetRequestId() *string
	SetSearchLibName(v string) *QuerySearchIndexResponseBody
	GetSearchLibName() *string
	SetSuccess(v string) *QuerySearchIndexResponseBody
	GetSuccess() *string
}

type QuerySearchIndexResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The state of the index. Valid values:
	//
	// 	- active: the index is enabled.
	//
	// 	- Deactive: the index is not enabled.
	//
	// example:
	//
	// Active
	IndexStatus *string `json:"IndexStatus,omitempty" xml:"IndexStatus,omitempty"`
	// The category of the index. Valid values:
	//
	// 	- mm: large visual model.
	//
	// 	- face: face recognition.
	//
	// 	- aiLabel: smart tagging.
	//
	// example:
	//
	// mm
	IndexType *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	// The total number of media assets.
	//
	// example:
	//
	// 12
	MediaTotal *string `json:"MediaTotal,omitempty" xml:"MediaTotal,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4E84BE44-58A7-****-****-FBEBEA16EF94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the search library.
	//
	// example:
	//
	// test1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySearchIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySearchIndexResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySearchIndexResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySearchIndexResponseBody) GetIndexStatus() *string {
	return s.IndexStatus
}

func (s *QuerySearchIndexResponseBody) GetIndexType() *string {
	return s.IndexType
}

func (s *QuerySearchIndexResponseBody) GetMediaTotal() *string {
	return s.MediaTotal
}

func (s *QuerySearchIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySearchIndexResponseBody) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *QuerySearchIndexResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *QuerySearchIndexResponseBody) SetCode(v string) *QuerySearchIndexResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySearchIndexResponseBody) SetIndexStatus(v string) *QuerySearchIndexResponseBody {
	s.IndexStatus = &v
	return s
}

func (s *QuerySearchIndexResponseBody) SetIndexType(v string) *QuerySearchIndexResponseBody {
	s.IndexType = &v
	return s
}

func (s *QuerySearchIndexResponseBody) SetMediaTotal(v string) *QuerySearchIndexResponseBody {
	s.MediaTotal = &v
	return s
}

func (s *QuerySearchIndexResponseBody) SetRequestId(v string) *QuerySearchIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySearchIndexResponseBody) SetSearchLibName(v string) *QuerySearchIndexResponseBody {
	s.SearchLibName = &v
	return s
}

func (s *QuerySearchIndexResponseBody) SetSuccess(v string) *QuerySearchIndexResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySearchIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
