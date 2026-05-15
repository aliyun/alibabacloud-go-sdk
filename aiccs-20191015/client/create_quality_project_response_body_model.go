// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateQualityProjectResponseBody
	GetCode() *string
	SetData(v *CreateQualityProjectResponseBodyData) *CreateQualityProjectResponseBody
	GetData() *CreateQualityProjectResponseBodyData
	SetMessage(v string) *CreateQualityProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateQualityProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateQualityProjectResponseBody
	GetSuccess() *bool
}

type CreateQualityProjectResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateQualityProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateQualityProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQualityProjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateQualityProjectResponseBody) GetData() *CreateQualityProjectResponseBodyData {
	return s.Data
}

func (s *CreateQualityProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateQualityProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQualityProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateQualityProjectResponseBody) SetCode(v string) *CreateQualityProjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateQualityProjectResponseBody) SetData(v *CreateQualityProjectResponseBodyData) *CreateQualityProjectResponseBody {
	s.Data = v
	return s
}

func (s *CreateQualityProjectResponseBody) SetMessage(v string) *CreateQualityProjectResponseBody {
	s.Message = &v
	return s
}

func (s *CreateQualityProjectResponseBody) SetRequestId(v string) *CreateQualityProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQualityProjectResponseBody) SetSuccess(v bool) *CreateQualityProjectResponseBody {
	s.Success = &v
	return s
}

func (s *CreateQualityProjectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateQualityProjectResponseBodyData struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProjectId  *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Version    *int32  `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateQualityProjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateQualityProjectResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateQualityProjectResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateQualityProjectResponseBodyData) GetVersion() *int32 {
	return s.Version
}

func (s *CreateQualityProjectResponseBodyData) SetInstanceId(v string) *CreateQualityProjectResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateQualityProjectResponseBodyData) SetProjectId(v int64) *CreateQualityProjectResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *CreateQualityProjectResponseBodyData) SetVersion(v int32) *CreateQualityProjectResponseBodyData {
	s.Version = &v
	return s
}

func (s *CreateQualityProjectResponseBodyData) Validate() error {
	return dara.Validate(s)
}
