// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMemoryResponseBody
	GetCode() *string
	SetData(v *GetMemoryResponseBodyData) *GetMemoryResponseBody
	GetData() *GetMemoryResponseBodyData
	SetRequestId(v string) *GetMemoryResponseBody
	GetRequestId() *string
}

type GetMemoryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                    `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetMemoryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 5A362ADD-51FC-5F4A-B858-D77F6EFAE2E6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMemoryResponseBody) GetData() *GetMemoryResponseBodyData {
	return s.Data
}

func (s *GetMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryResponseBody) SetCode(v string) *GetMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetMemoryResponseBody) SetData(v *GetMemoryResponseBodyData) *GetMemoryResponseBody {
	s.Data = v
	return s
}

func (s *GetMemoryResponseBody) SetRequestId(v string) *GetMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMemoryResponseBodyData struct {
	// example:
	//
	// my-cms-workspace
	CmsWorkspaceName *string `json:"cmsWorkspaceName,omitempty" xml:"cmsWorkspaceName,omitempty"`
	// example:
	//
	// 1735870116
	CreateTime *int32 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 365
	LongTtl *int32 `json:"longTtl,omitempty" xml:"longTtl,omitempty"`
	// example:
	//
	// test-memory
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 30
	ShortTtl *int32    `json:"shortTtl,omitempty" xml:"shortTtl,omitempty"`
	Strategy []*string `json:"strategy,omitempty" xml:"strategy,omitempty" type:"Repeated"`
}

func (s GetMemoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMemoryResponseBodyData) GetCmsWorkspaceName() *string {
	return s.CmsWorkspaceName
}

func (s *GetMemoryResponseBodyData) GetCreateTime() *int32 {
	return s.CreateTime
}

func (s *GetMemoryResponseBodyData) GetLongTtl() *int32 {
	return s.LongTtl
}

func (s *GetMemoryResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetMemoryResponseBodyData) GetShortTtl() *int32 {
	return s.ShortTtl
}

func (s *GetMemoryResponseBodyData) GetStrategy() []*string {
	return s.Strategy
}

func (s *GetMemoryResponseBodyData) SetCmsWorkspaceName(v string) *GetMemoryResponseBodyData {
	s.CmsWorkspaceName = &v
	return s
}

func (s *GetMemoryResponseBodyData) SetCreateTime(v int32) *GetMemoryResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMemoryResponseBodyData) SetLongTtl(v int32) *GetMemoryResponseBodyData {
	s.LongTtl = &v
	return s
}

func (s *GetMemoryResponseBodyData) SetName(v string) *GetMemoryResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMemoryResponseBodyData) SetShortTtl(v int32) *GetMemoryResponseBodyData {
	s.ShortTtl = &v
	return s
}

func (s *GetMemoryResponseBodyData) SetStrategy(v []*string) *GetMemoryResponseBodyData {
	s.Strategy = v
	return s
}

func (s *GetMemoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
