// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListInstanceStatusResponseBody
	GetRequestId() *string
	SetCode(v string) *ListInstanceStatusResponseBody
	GetCode() *string
	SetData(v []*ListInstanceStatusResponseBodyData) *ListInstanceStatusResponseBody
	GetData() []*ListInstanceStatusResponseBodyData
	SetMessage(v string) *ListInstanceStatusResponseBody
	GetMessage() *string
	SetTotal(v int64) *ListInstanceStatusResponseBody
	GetTotal() *int64
}

type ListInstanceStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListInstanceStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 218
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListInstanceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstanceStatusResponseBody) GetData() []*ListInstanceStatusResponseBodyData {
	return s.Data
}

func (s *ListInstanceStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstanceStatusResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListInstanceStatusResponseBody) SetRequestId(v string) *ListInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceStatusResponseBody) SetCode(v string) *ListInstanceStatusResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceStatusResponseBody) SetData(v []*ListInstanceStatusResponseBodyData) *ListInstanceStatusResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceStatusResponseBody) SetMessage(v string) *ListInstanceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceStatusResponseBody) SetTotal(v int64) *ListInstanceStatusResponseBody {
	s.Total = &v
	return s
}

func (s *ListInstanceStatusResponseBody) Validate() error {
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

type ListInstanceStatusResponseBodyData struct {
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	Region   *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListInstanceStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceStatusResponseBodyData) GetInstance() *string {
	return s.Instance
}

func (s *ListInstanceStatusResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *ListInstanceStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceStatusResponseBodyData) SetInstance(v string) *ListInstanceStatusResponseBodyData {
	s.Instance = &v
	return s
}

func (s *ListInstanceStatusResponseBodyData) SetRegion(v string) *ListInstanceStatusResponseBodyData {
	s.Region = &v
	return s
}

func (s *ListInstanceStatusResponseBodyData) SetStatus(v string) *ListInstanceStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListInstanceStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
