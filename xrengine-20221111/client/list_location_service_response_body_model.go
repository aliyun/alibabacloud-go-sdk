// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLocationServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLocationServiceResponseBody
	GetCode() *string
	SetCurrent(v int32) *ListLocationServiceResponseBody
	GetCurrent() *int32
	SetData(v []*ListLocationServiceResponseBodyData) *ListLocationServiceResponseBody
	GetData() []*ListLocationServiceResponseBodyData
	SetErrorName(v string) *ListLocationServiceResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *ListLocationServiceResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *ListLocationServiceResponseBody
	GetMessage() *string
	SetPages(v int32) *ListLocationServiceResponseBody
	GetPages() *int32
	SetRequestId(v string) *ListLocationServiceResponseBody
	GetRequestId() *string
	SetSize(v int32) *ListLocationServiceResponseBody
	GetSize() *int32
	SetSuccess(v bool) *ListLocationServiceResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListLocationServiceResponseBody
	GetTotal() *int32
}

type ListLocationServiceResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Current   *int32                                 `json:"Current,omitempty" xml:"Current,omitempty"`
	Data      []*ListLocationServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorName *string                                `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                                 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Pages     *int32                                 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size      *int32                                 `json:"Size,omitempty" xml:"Size,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32                                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListLocationServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLocationServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ListLocationServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLocationServiceResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *ListLocationServiceResponseBody) GetData() []*ListLocationServiceResponseBodyData {
	return s.Data
}

func (s *ListLocationServiceResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *ListLocationServiceResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListLocationServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLocationServiceResponseBody) GetPages() *int32 {
	return s.Pages
}

func (s *ListLocationServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLocationServiceResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *ListLocationServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLocationServiceResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListLocationServiceResponseBody) SetCode(v string) *ListLocationServiceResponseBody {
	s.Code = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetCurrent(v int32) *ListLocationServiceResponseBody {
	s.Current = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetData(v []*ListLocationServiceResponseBodyData) *ListLocationServiceResponseBody {
	s.Data = v
	return s
}

func (s *ListLocationServiceResponseBody) SetErrorName(v string) *ListLocationServiceResponseBody {
	s.ErrorName = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetHttpCode(v int32) *ListLocationServiceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetMessage(v string) *ListLocationServiceResponseBody {
	s.Message = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetPages(v int32) *ListLocationServiceResponseBody {
	s.Pages = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetRequestId(v string) *ListLocationServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetSize(v int32) *ListLocationServiceResponseBody {
	s.Size = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetSuccess(v bool) *ListLocationServiceResponseBody {
	s.Success = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetTotal(v int32) *ListLocationServiceResponseBody {
	s.Total = &v
	return s
}

func (s *ListLocationServiceResponseBody) Validate() error {
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

type ListLocationServiceResponseBodyData struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ExpireTime  *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Note        *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Qps         *int64  `json:"Qps,omitempty" xml:"Qps,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SvcState    *string `json:"SvcState,omitempty" xml:"SvcState,omitempty"`
}

func (s ListLocationServiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLocationServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLocationServiceResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *ListLocationServiceResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListLocationServiceResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListLocationServiceResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListLocationServiceResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListLocationServiceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListLocationServiceResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *ListLocationServiceResponseBodyData) GetQps() *int64 {
	return s.Qps
}

func (s *ListLocationServiceResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListLocationServiceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListLocationServiceResponseBodyData) GetSvcState() *string {
	return s.SvcState
}

func (s *ListLocationServiceResponseBodyData) SetAppId(v string) *ListLocationServiceResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetExpireTime(v string) *ListLocationServiceResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetGmtCreate(v string) *ListLocationServiceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetGmtModified(v string) *ListLocationServiceResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetId(v int64) *ListLocationServiceResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetName(v string) *ListLocationServiceResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetNote(v string) *ListLocationServiceResponseBodyData {
	s.Note = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetQps(v int64) *ListLocationServiceResponseBodyData {
	s.Qps = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetStartTime(v string) *ListLocationServiceResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetStatus(v string) *ListLocationServiceResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetSvcState(v string) *ListLocationServiceResponseBodyData {
	s.SvcState = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
