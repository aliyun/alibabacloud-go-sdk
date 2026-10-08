// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckConnectivityJobByJobIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCheckConnectivityJobByJobIdResponseBody
	GetCode() *string
	SetData(v *GetCheckConnectivityJobByJobIdResponseBodyData) *GetCheckConnectivityJobByJobIdResponseBody
	GetData() *GetCheckConnectivityJobByJobIdResponseBodyData
	SetHttpStatusCode(v int32) *GetCheckConnectivityJobByJobIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCheckConnectivityJobByJobIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCheckConnectivityJobByJobIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCheckConnectivityJobByJobIdResponseBody
	GetSuccess() *bool
}

type GetCheckConnectivityJobByJobIdResponseBody struct {
	// example:
	//
	// OK
	Code *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetCheckConnectivityJobByJobIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCheckConnectivityJobByJobIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCheckConnectivityJobByJobIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckConnectivityJobByJobIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCheckConnectivityJobByJobIdResponseBody) GetData() *GetCheckConnectivityJobByJobIdResponseBodyData {
	return s.Data
}

func (s *GetCheckConnectivityJobByJobIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCheckConnectivityJobByJobIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCheckConnectivityJobByJobIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCheckConnectivityJobByJobIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCheckConnectivityJobByJobIdResponseBody) SetCode(v string) *GetCheckConnectivityJobByJobIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetCheckConnectivityJobByJobIdResponseBody) SetData(v *GetCheckConnectivityJobByJobIdResponseBodyData) *GetCheckConnectivityJobByJobIdResponseBody {
	s.Data = v
	return s
}

func (s *GetCheckConnectivityJobByJobIdResponseBody) SetHttpStatusCode(v int32) *GetCheckConnectivityJobByJobIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCheckConnectivityJobByJobIdResponseBody) SetMessage(v string) *GetCheckConnectivityJobByJobIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetCheckConnectivityJobByJobIdResponseBody) SetRequestId(v string) *GetCheckConnectivityJobByJobIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCheckConnectivityJobByJobIdResponseBody) SetSuccess(v bool) *GetCheckConnectivityJobByJobIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetCheckConnectivityJobByJobIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCheckConnectivityJobByJobIdResponseBodyData struct {
	// example:
	//
	// 192
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// notFoundIp
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 123123
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// application/cluster
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 30001011
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// t_7572319950395080706_20251225_7572319950395080707
	VoldemortTaskId *string `json:"VoldemortTaskId,omitempty" xml:"VoldemortTaskId,omitempty"`
}

func (s GetCheckConnectivityJobByJobIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCheckConnectivityJobByJobIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCheckConnectivityJobByJobIdResponseBodyData) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *GetCheckConnectivityJobByJobIdResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetCheckConnectivityJobByJobIdResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *GetCheckConnectivityJobByJobIdResponseBodyData) GetJobType() *string {
	return s.JobType
}

func (s *GetCheckConnectivityJobByJobIdResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetCheckConnectivityJobByJobIdResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *GetCheckConnectivityJobByJobIdResponseBodyData) GetVoldemortTaskId() *string {
	return s.VoldemortTaskId
}

func (s *GetCheckConnectivityJobByJobIdResponseBodyData) SetDataSourceId(v string) *GetCheckConnectivityJobByJobIdResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *GetCheckConnectivityJobByJobIdResponseBodyData) SetErrorMsg(v string) *GetCheckConnectivityJobByJobIdResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *GetCheckConnectivityJobByJobIdResponseBodyData) SetJobId(v string) *GetCheckConnectivityJobByJobIdResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetCheckConnectivityJobByJobIdResponseBodyData) SetJobType(v string) *GetCheckConnectivityJobByJobIdResponseBodyData {
	s.JobType = &v
	return s
}

func (s *GetCheckConnectivityJobByJobIdResponseBodyData) SetStatus(v string) *GetCheckConnectivityJobByJobIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCheckConnectivityJobByJobIdResponseBodyData) SetTenantId(v string) *GetCheckConnectivityJobByJobIdResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetCheckConnectivityJobByJobIdResponseBodyData) SetVoldemortTaskId(v string) *GetCheckConnectivityJobByJobIdResponseBodyData {
	s.VoldemortTaskId = &v
	return s
}

func (s *GetCheckConnectivityJobByJobIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
