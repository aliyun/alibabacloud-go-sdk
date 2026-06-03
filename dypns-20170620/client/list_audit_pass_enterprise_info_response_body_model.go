// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuditPassEnterpriseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAuditPassEnterpriseInfoResponseBody
	GetCode() *string
	SetData(v []*ListAuditPassEnterpriseInfoResponseBodyData) *ListAuditPassEnterpriseInfoResponseBody
	GetData() []*ListAuditPassEnterpriseInfoResponseBodyData
	SetMessage(v string) *ListAuditPassEnterpriseInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAuditPassEnterpriseInfoResponseBody
	GetRequestId() *string
}

type ListAuditPassEnterpriseInfoResponseBody struct {
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListAuditPassEnterpriseInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAuditPassEnterpriseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuditPassEnterpriseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuditPassEnterpriseInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAuditPassEnterpriseInfoResponseBody) GetData() []*ListAuditPassEnterpriseInfoResponseBodyData {
	return s.Data
}

func (s *ListAuditPassEnterpriseInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAuditPassEnterpriseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuditPassEnterpriseInfoResponseBody) SetCode(v string) *ListAuditPassEnterpriseInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ListAuditPassEnterpriseInfoResponseBody) SetData(v []*ListAuditPassEnterpriseInfoResponseBodyData) *ListAuditPassEnterpriseInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListAuditPassEnterpriseInfoResponseBody) SetMessage(v string) *ListAuditPassEnterpriseInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ListAuditPassEnterpriseInfoResponseBody) SetRequestId(v string) *ListAuditPassEnterpriseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuditPassEnterpriseInfoResponseBody) Validate() error {
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

type ListAuditPassEnterpriseInfoResponseBodyData struct {
	EnterpriseId   *int64  `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	EnterpriseName *string `json:"EnterpriseName,omitempty" xml:"EnterpriseName,omitempty"`
}

func (s ListAuditPassEnterpriseInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAuditPassEnterpriseInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAuditPassEnterpriseInfoResponseBodyData) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListAuditPassEnterpriseInfoResponseBodyData) GetEnterpriseName() *string {
	return s.EnterpriseName
}

func (s *ListAuditPassEnterpriseInfoResponseBodyData) SetEnterpriseId(v int64) *ListAuditPassEnterpriseInfoResponseBodyData {
	s.EnterpriseId = &v
	return s
}

func (s *ListAuditPassEnterpriseInfoResponseBodyData) SetEnterpriseName(v string) *ListAuditPassEnterpriseInfoResponseBodyData {
	s.EnterpriseName = &v
	return s
}

func (s *ListAuditPassEnterpriseInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
