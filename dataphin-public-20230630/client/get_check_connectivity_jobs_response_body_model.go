// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckConnectivityJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCheckConnectivityJobsResponseBody
	GetCode() *string
	SetData(v []*GetCheckConnectivityJobsResponseBodyData) *GetCheckConnectivityJobsResponseBody
	GetData() []*GetCheckConnectivityJobsResponseBodyData
	SetHttpStatusCode(v int32) *GetCheckConnectivityJobsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCheckConnectivityJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCheckConnectivityJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCheckConnectivityJobsResponseBody
	GetSuccess() *bool
}

type GetCheckConnectivityJobsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// result
	Data []*GetCheckConnectivityJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCheckConnectivityJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCheckConnectivityJobsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckConnectivityJobsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCheckConnectivityJobsResponseBody) GetData() []*GetCheckConnectivityJobsResponseBodyData {
	return s.Data
}

func (s *GetCheckConnectivityJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCheckConnectivityJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCheckConnectivityJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCheckConnectivityJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCheckConnectivityJobsResponseBody) SetCode(v string) *GetCheckConnectivityJobsResponseBody {
	s.Code = &v
	return s
}

func (s *GetCheckConnectivityJobsResponseBody) SetData(v []*GetCheckConnectivityJobsResponseBodyData) *GetCheckConnectivityJobsResponseBody {
	s.Data = v
	return s
}

func (s *GetCheckConnectivityJobsResponseBody) SetHttpStatusCode(v int32) *GetCheckConnectivityJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCheckConnectivityJobsResponseBody) SetMessage(v string) *GetCheckConnectivityJobsResponseBody {
	s.Message = &v
	return s
}

func (s *GetCheckConnectivityJobsResponseBody) SetRequestId(v string) *GetCheckConnectivityJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCheckConnectivityJobsResponseBody) SetSuccess(v bool) *GetCheckConnectivityJobsResponseBody {
	s.Success = &v
	return s
}

func (s *GetCheckConnectivityJobsResponseBody) Validate() error {
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

type GetCheckConnectivityJobsResponseBodyData struct {
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

func (s GetCheckConnectivityJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCheckConnectivityJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCheckConnectivityJobsResponseBodyData) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *GetCheckConnectivityJobsResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetCheckConnectivityJobsResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *GetCheckConnectivityJobsResponseBodyData) GetJobType() *string {
	return s.JobType
}

func (s *GetCheckConnectivityJobsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetCheckConnectivityJobsResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *GetCheckConnectivityJobsResponseBodyData) GetVoldemortTaskId() *string {
	return s.VoldemortTaskId
}

func (s *GetCheckConnectivityJobsResponseBodyData) SetDataSourceId(v string) *GetCheckConnectivityJobsResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *GetCheckConnectivityJobsResponseBodyData) SetErrorMsg(v string) *GetCheckConnectivityJobsResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *GetCheckConnectivityJobsResponseBodyData) SetJobId(v string) *GetCheckConnectivityJobsResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetCheckConnectivityJobsResponseBodyData) SetJobType(v string) *GetCheckConnectivityJobsResponseBodyData {
	s.JobType = &v
	return s
}

func (s *GetCheckConnectivityJobsResponseBodyData) SetStatus(v string) *GetCheckConnectivityJobsResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCheckConnectivityJobsResponseBodyData) SetTenantId(v string) *GetCheckConnectivityJobsResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetCheckConnectivityJobsResponseBodyData) SetVoldemortTaskId(v string) *GetCheckConnectivityJobsResponseBodyData {
	s.VoldemortTaskId = &v
	return s
}

func (s *GetCheckConnectivityJobsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
