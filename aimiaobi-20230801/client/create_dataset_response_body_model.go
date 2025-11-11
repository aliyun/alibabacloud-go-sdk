// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDatasetResponseBody
	GetCode() *string
	SetData(v *CreateDatasetResponseBodyData) *CreateDatasetResponseBody
	GetData() *CreateDatasetResponseBodyData
	SetHttpStatusCode(v int32) *CreateDatasetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDatasetResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDatasetResponseBody
	GetSuccess() *bool
}

type CreateDatasetResponseBody struct {
	// example:
	//
	// NoData
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateDatasetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDatasetResponseBody) GetData() *CreateDatasetResponseBodyData {
	return s.Data
}

func (s *CreateDatasetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDatasetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDatasetResponseBody) SetCode(v string) *CreateDatasetResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDatasetResponseBody) SetData(v *CreateDatasetResponseBodyData) *CreateDatasetResponseBody {
	s.Data = v
	return s
}

func (s *CreateDatasetResponseBody) SetHttpStatusCode(v int32) *CreateDatasetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDatasetResponseBody) SetMessage(v string) *CreateDatasetResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDatasetResponseBody) SetRequestId(v string) *CreateDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetResponseBody) SetSuccess(v bool) *CreateDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDatasetResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDatasetResponseBodyData struct {
	// example:
	//
	// 2024-11-12 21:46:24
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// xxx
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// xxx
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// xxx
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// CustomSemanticSearch
	DatasetType *string `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	// example:
	//
	// 1
	SearchDatasetEnable *int32 `json:"SearchDatasetEnable,omitempty" xml:"SearchDatasetEnable,omitempty"`
}

func (s CreateDatasetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateDatasetResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *CreateDatasetResponseBodyData) GetDatasetDescription() *string {
	return s.DatasetDescription
}

func (s *CreateDatasetResponseBodyData) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *CreateDatasetResponseBodyData) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateDatasetResponseBodyData) GetDatasetType() *string {
	return s.DatasetType
}

func (s *CreateDatasetResponseBodyData) GetSearchDatasetEnable() *int32 {
	return s.SearchDatasetEnable
}

func (s *CreateDatasetResponseBodyData) SetCreateTime(v string) *CreateDatasetResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateDatasetResponseBodyData) SetCreateUser(v string) *CreateDatasetResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *CreateDatasetResponseBodyData) SetDatasetDescription(v string) *CreateDatasetResponseBodyData {
	s.DatasetDescription = &v
	return s
}

func (s *CreateDatasetResponseBodyData) SetDatasetId(v int64) *CreateDatasetResponseBodyData {
	s.DatasetId = &v
	return s
}

func (s *CreateDatasetResponseBodyData) SetDatasetName(v string) *CreateDatasetResponseBodyData {
	s.DatasetName = &v
	return s
}

func (s *CreateDatasetResponseBodyData) SetDatasetType(v string) *CreateDatasetResponseBodyData {
	s.DatasetType = &v
	return s
}

func (s *CreateDatasetResponseBodyData) SetSearchDatasetEnable(v int32) *CreateDatasetResponseBodyData {
	s.SearchDatasetEnable = &v
	return s
}

func (s *CreateDatasetResponseBodyData) Validate() error {
	return dara.Validate(s)
}
