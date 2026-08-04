// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBindsByOuterIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryBindsByOuterIdResponseBody
	GetCode() *string
	SetData(v []*QueryBindsByOuterIdResponseBodyData) *QueryBindsByOuterIdResponseBody
	GetData() []*QueryBindsByOuterIdResponseBodyData
	SetHttpCode(v string) *QueryBindsByOuterIdResponseBody
	GetHttpCode() *string
	SetMessage(v string) *QueryBindsByOuterIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryBindsByOuterIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryBindsByOuterIdResponseBody
	GetSuccess() *bool
}

type QueryBindsByOuterIdResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*QueryBindsByOuterIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpCode  *string                                `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryBindsByOuterIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryBindsByOuterIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBindsByOuterIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryBindsByOuterIdResponseBody) GetData() []*QueryBindsByOuterIdResponseBodyData {
	return s.Data
}

func (s *QueryBindsByOuterIdResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *QueryBindsByOuterIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryBindsByOuterIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryBindsByOuterIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryBindsByOuterIdResponseBody) SetCode(v string) *QueryBindsByOuterIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBindsByOuterIdResponseBody) SetData(v []*QueryBindsByOuterIdResponseBodyData) *QueryBindsByOuterIdResponseBody {
	s.Data = v
	return s
}

func (s *QueryBindsByOuterIdResponseBody) SetHttpCode(v string) *QueryBindsByOuterIdResponseBody {
	s.HttpCode = &v
	return s
}

func (s *QueryBindsByOuterIdResponseBody) SetMessage(v string) *QueryBindsByOuterIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBindsByOuterIdResponseBody) SetRequestId(v string) *QueryBindsByOuterIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBindsByOuterIdResponseBody) SetSuccess(v bool) *QueryBindsByOuterIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryBindsByOuterIdResponseBody) Validate() error {
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

type QueryBindsByOuterIdResponseBodyData struct {
	BindData     map[string]interface{} `json:"BindData,omitempty" xml:"BindData,omitempty"`
	MinorOuterId *string                `json:"MinorOuterId,omitempty" xml:"MinorOuterId,omitempty"`
	OuterId      *string                `json:"OuterId,omitempty" xml:"OuterId,omitempty"`
	Pk           *string                `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Status       *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId     *string                `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryBindsByOuterIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryBindsByOuterIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBindsByOuterIdResponseBodyData) GetBindData() map[string]interface{} {
	return s.BindData
}

func (s *QueryBindsByOuterIdResponseBodyData) GetMinorOuterId() *string {
	return s.MinorOuterId
}

func (s *QueryBindsByOuterIdResponseBodyData) GetOuterId() *string {
	return s.OuterId
}

func (s *QueryBindsByOuterIdResponseBodyData) GetPk() *string {
	return s.Pk
}

func (s *QueryBindsByOuterIdResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryBindsByOuterIdResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryBindsByOuterIdResponseBodyData) SetBindData(v map[string]interface{}) *QueryBindsByOuterIdResponseBodyData {
	s.BindData = v
	return s
}

func (s *QueryBindsByOuterIdResponseBodyData) SetMinorOuterId(v string) *QueryBindsByOuterIdResponseBodyData {
	s.MinorOuterId = &v
	return s
}

func (s *QueryBindsByOuterIdResponseBodyData) SetOuterId(v string) *QueryBindsByOuterIdResponseBodyData {
	s.OuterId = &v
	return s
}

func (s *QueryBindsByOuterIdResponseBodyData) SetPk(v string) *QueryBindsByOuterIdResponseBodyData {
	s.Pk = &v
	return s
}

func (s *QueryBindsByOuterIdResponseBodyData) SetStatus(v string) *QueryBindsByOuterIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryBindsByOuterIdResponseBodyData) SetTenantId(v string) *QueryBindsByOuterIdResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *QueryBindsByOuterIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
