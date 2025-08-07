// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErrorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v map[string]interface{}) *GetErrorsResponseBody
	GetArgs() map[string]interface{}
	SetErrorCode(v int32) *GetErrorsResponseBody
	GetErrorCode() *int32
	SetMessage(v string) *GetErrorsResponseBody
	GetMessage() *string
	SetModel(v *GetErrorsResponseBodyModel) *GetErrorsResponseBody
	GetModel() *GetErrorsResponseBodyModel
	SetRequestId(v string) *GetErrorsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetErrorsResponseBody
	GetSuccess() *bool
}

type GetErrorsResponseBody struct {
	// Args
	Args map[string]interface{} `json:"Args,omitempty" xml:"Args,omitempty"`
	// example:
	//
	// 500
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// internal error
	Message *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *GetErrorsResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// RequestId
	//
	// example:
	//
	// B3AD0FE4-36EF-1641-90B1-77618166F2ff
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetErrorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetErrorsResponseBody) GoString() string {
	return s.String()
}

func (s *GetErrorsResponseBody) GetArgs() map[string]interface{} {
	return s.Args
}

func (s *GetErrorsResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetErrorsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetErrorsResponseBody) GetModel() *GetErrorsResponseBodyModel {
	return s.Model
}

func (s *GetErrorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetErrorsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetErrorsResponseBody) SetArgs(v map[string]interface{}) *GetErrorsResponseBody {
	s.Args = v
	return s
}

func (s *GetErrorsResponseBody) SetErrorCode(v int32) *GetErrorsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetErrorsResponseBody) SetMessage(v string) *GetErrorsResponseBody {
	s.Message = &v
	return s
}

func (s *GetErrorsResponseBody) SetModel(v *GetErrorsResponseBodyModel) *GetErrorsResponseBody {
	s.Model = v
	return s
}

func (s *GetErrorsResponseBody) SetRequestId(v string) *GetErrorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetErrorsResponseBody) SetSuccess(v bool) *GetErrorsResponseBody {
	s.Success = &v
	return s
}

func (s *GetErrorsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetErrorsResponseBodyModel struct {
	Items []*GetErrorsResponseBodyModelItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	Pages *int32 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetErrorsResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s GetErrorsResponseBodyModel) GoString() string {
	return s.String()
}

func (s *GetErrorsResponseBodyModel) GetItems() []*GetErrorsResponseBodyModelItems {
	return s.Items
}

func (s *GetErrorsResponseBodyModel) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetErrorsResponseBodyModel) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetErrorsResponseBodyModel) GetPages() *int32 {
	return s.Pages
}

func (s *GetErrorsResponseBodyModel) GetTotal() *int64 {
	return s.Total
}

func (s *GetErrorsResponseBodyModel) SetItems(v []*GetErrorsResponseBodyModelItems) *GetErrorsResponseBodyModel {
	s.Items = v
	return s
}

func (s *GetErrorsResponseBodyModel) SetPageNum(v int32) *GetErrorsResponseBodyModel {
	s.PageNum = &v
	return s
}

func (s *GetErrorsResponseBodyModel) SetPageSize(v int32) *GetErrorsResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *GetErrorsResponseBodyModel) SetPages(v int32) *GetErrorsResponseBodyModel {
	s.Pages = &v
	return s
}

func (s *GetErrorsResponseBodyModel) SetTotal(v int64) *GetErrorsResponseBodyModel {
	s.Total = &v
	return s
}

func (s *GetErrorsResponseBodyModel) Validate() error {
	return dara.Validate(s)
}

type GetErrorsResponseBodyModelItems struct {
	// example:
	//
	// 1740488561065
	ClientTime *int64 `json:"ClientTime,omitempty" xml:"ClientTime,omitempty"`
	// example:
	//
	// RANDOM-1729634758587856312DEVICE
	Did        *string `json:"Did,omitempty" xml:"Did,omitempty"`
	DigestHash *string `json:"DigestHash,omitempty" xml:"DigestHash,omitempty"`
	// Utdid
	//
	// example:
	//
	// RANDOM-1729634758587856312DEVICE
	Utdid *string `json:"Utdid,omitempty" xml:"Utdid,omitempty"`
	// example:
	//
	// 9634758587856312DEV
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetErrorsResponseBodyModelItems) String() string {
	return dara.Prettify(s)
}

func (s GetErrorsResponseBodyModelItems) GoString() string {
	return s.String()
}

func (s *GetErrorsResponseBodyModelItems) GetClientTime() *int64 {
	return s.ClientTime
}

func (s *GetErrorsResponseBodyModelItems) GetDid() *string {
	return s.Did
}

func (s *GetErrorsResponseBodyModelItems) GetDigestHash() *string {
	return s.DigestHash
}

func (s *GetErrorsResponseBodyModelItems) GetUtdid() *string {
	return s.Utdid
}

func (s *GetErrorsResponseBodyModelItems) GetUuid() *string {
	return s.Uuid
}

func (s *GetErrorsResponseBodyModelItems) SetClientTime(v int64) *GetErrorsResponseBodyModelItems {
	s.ClientTime = &v
	return s
}

func (s *GetErrorsResponseBodyModelItems) SetDid(v string) *GetErrorsResponseBodyModelItems {
	s.Did = &v
	return s
}

func (s *GetErrorsResponseBodyModelItems) SetDigestHash(v string) *GetErrorsResponseBodyModelItems {
	s.DigestHash = &v
	return s
}

func (s *GetErrorsResponseBodyModelItems) SetUtdid(v string) *GetErrorsResponseBodyModelItems {
	s.Utdid = &v
	return s
}

func (s *GetErrorsResponseBodyModelItems) SetUuid(v string) *GetErrorsResponseBodyModelItems {
	s.Uuid = &v
	return s
}

func (s *GetErrorsResponseBodyModelItems) Validate() error {
	return dara.Validate(s)
}
