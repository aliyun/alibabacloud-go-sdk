// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityClassifyCatalogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSecurityClassifyCatalogResponseBody
	GetCode() *string
	SetData(v *DeleteSecurityClassifyCatalogResponseBodyData) *DeleteSecurityClassifyCatalogResponseBody
	GetData() *DeleteSecurityClassifyCatalogResponseBodyData
	SetHttpStatusCode(v int32) *DeleteSecurityClassifyCatalogResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteSecurityClassifyCatalogResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSecurityClassifyCatalogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSecurityClassifyCatalogResponseBody
	GetSuccess() *bool
}

type DeleteSecurityClassifyCatalogResponseBody struct {
	// example:
	//
	// OK
	Code *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteSecurityClassifyCatalogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSecurityClassifyCatalogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityClassifyCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityClassifyCatalogResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSecurityClassifyCatalogResponseBody) GetData() *DeleteSecurityClassifyCatalogResponseBodyData {
	return s.Data
}

func (s *DeleteSecurityClassifyCatalogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteSecurityClassifyCatalogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSecurityClassifyCatalogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecurityClassifyCatalogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSecurityClassifyCatalogResponseBody) SetCode(v string) *DeleteSecurityClassifyCatalogResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSecurityClassifyCatalogResponseBody) SetData(v *DeleteSecurityClassifyCatalogResponseBodyData) *DeleteSecurityClassifyCatalogResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSecurityClassifyCatalogResponseBody) SetHttpStatusCode(v int32) *DeleteSecurityClassifyCatalogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteSecurityClassifyCatalogResponseBody) SetMessage(v string) *DeleteSecurityClassifyCatalogResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSecurityClassifyCatalogResponseBody) SetRequestId(v string) *DeleteSecurityClassifyCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecurityClassifyCatalogResponseBody) SetSuccess(v bool) *DeleteSecurityClassifyCatalogResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSecurityClassifyCatalogResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSecurityClassifyCatalogResponseBodyData struct {
	ChildCatalogFullPathList []*string `json:"ChildCatalogFullPathList,omitempty" xml:"ChildCatalogFullPathList,omitempty" type:"Repeated"`
	ClassifyIdList           []*int64  `json:"ClassifyIdList,omitempty" xml:"ClassifyIdList,omitempty" type:"Repeated"`
	Success                  *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSecurityClassifyCatalogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityClassifyCatalogResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteSecurityClassifyCatalogResponseBodyData) GetChildCatalogFullPathList() []*string {
	return s.ChildCatalogFullPathList
}

func (s *DeleteSecurityClassifyCatalogResponseBodyData) GetClassifyIdList() []*int64 {
	return s.ClassifyIdList
}

func (s *DeleteSecurityClassifyCatalogResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSecurityClassifyCatalogResponseBodyData) SetChildCatalogFullPathList(v []*string) *DeleteSecurityClassifyCatalogResponseBodyData {
	s.ChildCatalogFullPathList = v
	return s
}

func (s *DeleteSecurityClassifyCatalogResponseBodyData) SetClassifyIdList(v []*int64) *DeleteSecurityClassifyCatalogResponseBodyData {
	s.ClassifyIdList = v
	return s
}

func (s *DeleteSecurityClassifyCatalogResponseBodyData) SetSuccess(v bool) *DeleteSecurityClassifyCatalogResponseBodyData {
	s.Success = &v
	return s
}

func (s *DeleteSecurityClassifyCatalogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
