// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertDestinationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAlertDestinationResponseBody
	GetCode() *string
	SetData(v *GetAlertDestinationResponseBodyData) *GetAlertDestinationResponseBody
	GetData() *GetAlertDestinationResponseBodyData
	SetMessage(v string) *GetAlertDestinationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAlertDestinationResponseBody
	GetRequestId() *string
}

type GetAlertDestinationResponseBody struct {
	// example:
	//
	// Success or Sysom.ServerError
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetAlertDestinationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAlertDestinationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlertDestinationResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlertDestinationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAlertDestinationResponseBody) GetData() *GetAlertDestinationResponseBodyData {
	return s.Data
}

func (s *GetAlertDestinationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAlertDestinationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlertDestinationResponseBody) SetCode(v string) *GetAlertDestinationResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlertDestinationResponseBody) SetData(v *GetAlertDestinationResponseBodyData) *GetAlertDestinationResponseBody {
	s.Data = v
	return s
}

func (s *GetAlertDestinationResponseBody) SetMessage(v string) *GetAlertDestinationResponseBody {
	s.Message = &v
	return s
}

func (s *GetAlertDestinationResponseBody) SetRequestId(v string) *GetAlertDestinationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlertDestinationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlertDestinationResponseBodyData struct {
	// example:
	//
	// 1751520976660
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 1
	Id *int32 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// name1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// {
	//
	//     "webhook":"",
	//
	//     "sec":"",
	//
	// }
	Params interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// example:
	//
	// console
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// dingtalk
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
	// example:
	//
	// 1234123412352311
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// example:
	//
	// 1751254826285
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s GetAlertDestinationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAlertDestinationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAlertDestinationResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetAlertDestinationResponseBodyData) GetId() *int32 {
	return s.Id
}

func (s *GetAlertDestinationResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetAlertDestinationResponseBodyData) GetParams() interface{} {
	return s.Params
}

func (s *GetAlertDestinationResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *GetAlertDestinationResponseBodyData) GetTarget() *string {
	return s.Target
}

func (s *GetAlertDestinationResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *GetAlertDestinationResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetAlertDestinationResponseBodyData) SetCreatedAt(v string) *GetAlertDestinationResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetAlertDestinationResponseBodyData) SetId(v int32) *GetAlertDestinationResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetAlertDestinationResponseBodyData) SetName(v string) *GetAlertDestinationResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAlertDestinationResponseBodyData) SetParams(v interface{}) *GetAlertDestinationResponseBodyData {
	s.Params = v
	return s
}

func (s *GetAlertDestinationResponseBodyData) SetSource(v string) *GetAlertDestinationResponseBodyData {
	s.Source = &v
	return s
}

func (s *GetAlertDestinationResponseBodyData) SetTarget(v string) *GetAlertDestinationResponseBodyData {
	s.Target = &v
	return s
}

func (s *GetAlertDestinationResponseBodyData) SetUid(v string) *GetAlertDestinationResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetAlertDestinationResponseBodyData) SetUpdatedAt(v string) *GetAlertDestinationResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetAlertDestinationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
