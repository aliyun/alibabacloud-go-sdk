// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetServiceAccessResponseBody
	GetRequestId() *string
	SetServiceAccessInfo(v *GetServiceAccessResponseBodyServiceAccessInfo) *GetServiceAccessResponseBody
	GetServiceAccessInfo() *GetServiceAccessResponseBodyServiceAccessInfo
}

type GetServiceAccessResponseBody struct {
	RequestId         *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceAccessInfo *GetServiceAccessResponseBodyServiceAccessInfo `json:"ServiceAccessInfo,omitempty" xml:"ServiceAccessInfo,omitempty" type:"Struct"`
}

func (s GetServiceAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceAccessResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceAccessResponseBody) GetServiceAccessInfo() *GetServiceAccessResponseBodyServiceAccessInfo {
	return s.ServiceAccessInfo
}

func (s *GetServiceAccessResponseBody) SetRequestId(v string) *GetServiceAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceAccessResponseBody) SetServiceAccessInfo(v *GetServiceAccessResponseBodyServiceAccessInfo) *GetServiceAccessResponseBody {
	s.ServiceAccessInfo = v
	return s
}

func (s *GetServiceAccessResponseBody) Validate() error {
	if s.ServiceAccessInfo != nil {
		if err := s.ServiceAccessInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceAccessResponseBodyServiceAccessInfo struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetServiceAccessResponseBodyServiceAccessInfo) String() string {
	return dara.Prettify(s)
}

func (s GetServiceAccessResponseBodyServiceAccessInfo) GoString() string {
	return s.String()
}

func (s *GetServiceAccessResponseBodyServiceAccessInfo) GetStatus() *string {
	return s.Status
}

func (s *GetServiceAccessResponseBodyServiceAccessInfo) SetStatus(v string) *GetServiceAccessResponseBodyServiceAccessInfo {
	s.Status = &v
	return s
}

func (s *GetServiceAccessResponseBodyServiceAccessInfo) Validate() error {
	return dara.Validate(s)
}
