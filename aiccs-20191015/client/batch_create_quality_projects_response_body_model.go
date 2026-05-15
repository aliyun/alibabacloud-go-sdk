// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateQualityProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchCreateQualityProjectsResponseBody
	GetCode() *string
	SetData(v []*BatchCreateQualityProjectsResponseBodyData) *BatchCreateQualityProjectsResponseBody
	GetData() []*BatchCreateQualityProjectsResponseBodyData
	SetMessage(v string) *BatchCreateQualityProjectsResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchCreateQualityProjectsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchCreateQualityProjectsResponseBody
	GetSuccess() *bool
}

type BatchCreateQualityProjectsResponseBody struct {
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*BatchCreateQualityProjectsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchCreateQualityProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateQualityProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateQualityProjectsResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchCreateQualityProjectsResponseBody) GetData() []*BatchCreateQualityProjectsResponseBodyData {
	return s.Data
}

func (s *BatchCreateQualityProjectsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchCreateQualityProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCreateQualityProjectsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchCreateQualityProjectsResponseBody) SetCode(v string) *BatchCreateQualityProjectsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchCreateQualityProjectsResponseBody) SetData(v []*BatchCreateQualityProjectsResponseBodyData) *BatchCreateQualityProjectsResponseBody {
	s.Data = v
	return s
}

func (s *BatchCreateQualityProjectsResponseBody) SetMessage(v string) *BatchCreateQualityProjectsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchCreateQualityProjectsResponseBody) SetRequestId(v string) *BatchCreateQualityProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateQualityProjectsResponseBody) SetSuccess(v bool) *BatchCreateQualityProjectsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchCreateQualityProjectsResponseBody) Validate() error {
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

type BatchCreateQualityProjectsResponseBodyData struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProjectId  *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Version    *int32  `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s BatchCreateQualityProjectsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateQualityProjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchCreateQualityProjectsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BatchCreateQualityProjectsResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *BatchCreateQualityProjectsResponseBodyData) GetVersion() *int32 {
	return s.Version
}

func (s *BatchCreateQualityProjectsResponseBodyData) SetInstanceId(v string) *BatchCreateQualityProjectsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *BatchCreateQualityProjectsResponseBodyData) SetProjectId(v int64) *BatchCreateQualityProjectsResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *BatchCreateQualityProjectsResponseBodyData) SetVersion(v int32) *BatchCreateQualityProjectsResponseBodyData {
	s.Version = &v
	return s
}

func (s *BatchCreateQualityProjectsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
