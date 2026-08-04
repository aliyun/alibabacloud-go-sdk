// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBindsByPkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryBindsByPkResponseBody
	GetCode() *string
	SetData(v []*QueryBindsByPkResponseBodyData) *QueryBindsByPkResponseBody
	GetData() []*QueryBindsByPkResponseBodyData
	SetHttpCode(v string) *QueryBindsByPkResponseBody
	GetHttpCode() *string
	SetMessage(v string) *QueryBindsByPkResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryBindsByPkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryBindsByPkResponseBody
	GetSuccess() *bool
}

type QueryBindsByPkResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*QueryBindsByPkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpCode  *string                           `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryBindsByPkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryBindsByPkResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBindsByPkResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryBindsByPkResponseBody) GetData() []*QueryBindsByPkResponseBodyData {
	return s.Data
}

func (s *QueryBindsByPkResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *QueryBindsByPkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryBindsByPkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryBindsByPkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryBindsByPkResponseBody) SetCode(v string) *QueryBindsByPkResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBindsByPkResponseBody) SetData(v []*QueryBindsByPkResponseBodyData) *QueryBindsByPkResponseBody {
	s.Data = v
	return s
}

func (s *QueryBindsByPkResponseBody) SetHttpCode(v string) *QueryBindsByPkResponseBody {
	s.HttpCode = &v
	return s
}

func (s *QueryBindsByPkResponseBody) SetMessage(v string) *QueryBindsByPkResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBindsByPkResponseBody) SetRequestId(v string) *QueryBindsByPkResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBindsByPkResponseBody) SetSuccess(v bool) *QueryBindsByPkResponseBody {
	s.Success = &v
	return s
}

func (s *QueryBindsByPkResponseBody) Validate() error {
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

type QueryBindsByPkResponseBodyData struct {
	BindData     map[string]interface{} `json:"BindData,omitempty" xml:"BindData,omitempty"`
	MinorOuterId *string                `json:"MinorOuterId,omitempty" xml:"MinorOuterId,omitempty"`
	OuterId      *string                `json:"OuterId,omitempty" xml:"OuterId,omitempty"`
	Pk           *string                `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Status       *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId     *string                `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryBindsByPkResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryBindsByPkResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBindsByPkResponseBodyData) GetBindData() map[string]interface{} {
	return s.BindData
}

func (s *QueryBindsByPkResponseBodyData) GetMinorOuterId() *string {
	return s.MinorOuterId
}

func (s *QueryBindsByPkResponseBodyData) GetOuterId() *string {
	return s.OuterId
}

func (s *QueryBindsByPkResponseBodyData) GetPk() *string {
	return s.Pk
}

func (s *QueryBindsByPkResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryBindsByPkResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryBindsByPkResponseBodyData) SetBindData(v map[string]interface{}) *QueryBindsByPkResponseBodyData {
	s.BindData = v
	return s
}

func (s *QueryBindsByPkResponseBodyData) SetMinorOuterId(v string) *QueryBindsByPkResponseBodyData {
	s.MinorOuterId = &v
	return s
}

func (s *QueryBindsByPkResponseBodyData) SetOuterId(v string) *QueryBindsByPkResponseBodyData {
	s.OuterId = &v
	return s
}

func (s *QueryBindsByPkResponseBodyData) SetPk(v string) *QueryBindsByPkResponseBodyData {
	s.Pk = &v
	return s
}

func (s *QueryBindsByPkResponseBodyData) SetStatus(v string) *QueryBindsByPkResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryBindsByPkResponseBodyData) SetTenantId(v string) *QueryBindsByPkResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *QueryBindsByPkResponseBodyData) Validate() error {
	return dara.Validate(s)
}
