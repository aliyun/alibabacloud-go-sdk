// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySearchLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySearchLibResponseBody
	GetCode() *string
	SetIndexInfo(v []*QuerySearchLibResponseBodyIndexInfo) *QuerySearchLibResponseBody
	GetIndexInfo() []*QuerySearchLibResponseBodyIndexInfo
	SetRequestId(v string) *QuerySearchLibResponseBody
	GetRequestId() *string
	SetSearchLibConfig(v string) *QuerySearchLibResponseBody
	GetSearchLibConfig() *string
	SetSearchLibName(v string) *QuerySearchLibResponseBody
	GetSearchLibName() *string
	SetStatus(v string) *QuerySearchLibResponseBody
	GetStatus() *string
	SetSuccess(v string) *QuerySearchLibResponseBody
	GetSuccess() *string
}

type QuerySearchLibResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	IndexInfo []*QuerySearchLibResponseBodyIndexInfo `json:"IndexInfo,omitempty" xml:"IndexInfo,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SearchLibConfig *string `json:"SearchLibConfig,omitempty" xml:"SearchLibConfig,omitempty"`
	// The name of the search library.
	//
	// example:
	//
	// test1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
	// The status of the search library.
	//
	// Valid values:
	//
	// 	- normal
	//
	// 	- deleting
	//
	// 	- deleteFail
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the call was successful. Valid values:
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

func (s QuerySearchLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySearchLibResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySearchLibResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySearchLibResponseBody) GetIndexInfo() []*QuerySearchLibResponseBodyIndexInfo {
	return s.IndexInfo
}

func (s *QuerySearchLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySearchLibResponseBody) GetSearchLibConfig() *string {
	return s.SearchLibConfig
}

func (s *QuerySearchLibResponseBody) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *QuerySearchLibResponseBody) GetStatus() *string {
	return s.Status
}

func (s *QuerySearchLibResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *QuerySearchLibResponseBody) SetCode(v string) *QuerySearchLibResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySearchLibResponseBody) SetIndexInfo(v []*QuerySearchLibResponseBodyIndexInfo) *QuerySearchLibResponseBody {
	s.IndexInfo = v
	return s
}

func (s *QuerySearchLibResponseBody) SetRequestId(v string) *QuerySearchLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySearchLibResponseBody) SetSearchLibConfig(v string) *QuerySearchLibResponseBody {
	s.SearchLibConfig = &v
	return s
}

func (s *QuerySearchLibResponseBody) SetSearchLibName(v string) *QuerySearchLibResponseBody {
	s.SearchLibName = &v
	return s
}

func (s *QuerySearchLibResponseBody) SetStatus(v string) *QuerySearchLibResponseBody {
	s.Status = &v
	return s
}

func (s *QuerySearchLibResponseBody) SetSuccess(v string) *QuerySearchLibResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySearchLibResponseBody) Validate() error {
	if s.IndexInfo != nil {
		for _, item := range s.IndexInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySearchLibResponseBodyIndexInfo struct {
	IndexReadiness *string `json:"IndexReadiness,omitempty" xml:"IndexReadiness,omitempty"`
	IndexStatus    *string `json:"IndexStatus,omitempty" xml:"IndexStatus,omitempty"`
	IndexType      *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
}

func (s QuerySearchLibResponseBodyIndexInfo) String() string {
	return dara.Prettify(s)
}

func (s QuerySearchLibResponseBodyIndexInfo) GoString() string {
	return s.String()
}

func (s *QuerySearchLibResponseBodyIndexInfo) GetIndexReadiness() *string {
	return s.IndexReadiness
}

func (s *QuerySearchLibResponseBodyIndexInfo) GetIndexStatus() *string {
	return s.IndexStatus
}

func (s *QuerySearchLibResponseBodyIndexInfo) GetIndexType() *string {
	return s.IndexType
}

func (s *QuerySearchLibResponseBodyIndexInfo) SetIndexReadiness(v string) *QuerySearchLibResponseBodyIndexInfo {
	s.IndexReadiness = &v
	return s
}

func (s *QuerySearchLibResponseBodyIndexInfo) SetIndexStatus(v string) *QuerySearchLibResponseBodyIndexInfo {
	s.IndexStatus = &v
	return s
}

func (s *QuerySearchLibResponseBodyIndexInfo) SetIndexType(v string) *QuerySearchLibResponseBodyIndexInfo {
	s.IndexType = &v
	return s
}

func (s *QuerySearchLibResponseBodyIndexInfo) Validate() error {
	return dara.Validate(s)
}
