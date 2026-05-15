// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityProjectLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityProjectLogResponseBody
	GetCode() *string
	SetData(v []*GetQualityProjectLogResponseBodyData) *GetQualityProjectLogResponseBody
	GetData() []*GetQualityProjectLogResponseBodyData
	SetMessage(v string) *GetQualityProjectLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQualityProjectLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityProjectLogResponseBody
	GetSuccess() *bool
}

type GetQualityProjectLogResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetQualityProjectLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityProjectLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityProjectLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityProjectLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityProjectLogResponseBody) GetData() []*GetQualityProjectLogResponseBodyData {
	return s.Data
}

func (s *GetQualityProjectLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityProjectLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityProjectLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityProjectLogResponseBody) SetCode(v string) *GetQualityProjectLogResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityProjectLogResponseBody) SetData(v []*GetQualityProjectLogResponseBodyData) *GetQualityProjectLogResponseBody {
	s.Data = v
	return s
}

func (s *GetQualityProjectLogResponseBody) SetMessage(v string) *GetQualityProjectLogResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityProjectLogResponseBody) SetRequestId(v string) *GetQualityProjectLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityProjectLogResponseBody) SetSuccess(v bool) *GetQualityProjectLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityProjectLogResponseBody) Validate() error {
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

type GetQualityProjectLogResponseBodyData struct {
	ActionData        *string `json:"ActionData,omitempty" xml:"ActionData,omitempty"`
	ActionTime        *string `json:"ActionTime,omitempty" xml:"ActionTime,omitempty"`
	ActionType        *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	ProjectCreateTime *string `json:"ProjectCreateTime,omitempty" xml:"ProjectCreateTime,omitempty"`
	ProjectId         *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetQualityProjectLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQualityProjectLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityProjectLogResponseBodyData) GetActionData() *string {
	return s.ActionData
}

func (s *GetQualityProjectLogResponseBodyData) GetActionTime() *string {
	return s.ActionTime
}

func (s *GetQualityProjectLogResponseBodyData) GetActionType() *string {
	return s.ActionType
}

func (s *GetQualityProjectLogResponseBodyData) GetProjectCreateTime() *string {
	return s.ProjectCreateTime
}

func (s *GetQualityProjectLogResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetQualityProjectLogResponseBodyData) SetActionData(v string) *GetQualityProjectLogResponseBodyData {
	s.ActionData = &v
	return s
}

func (s *GetQualityProjectLogResponseBodyData) SetActionTime(v string) *GetQualityProjectLogResponseBodyData {
	s.ActionTime = &v
	return s
}

func (s *GetQualityProjectLogResponseBodyData) SetActionType(v string) *GetQualityProjectLogResponseBodyData {
	s.ActionType = &v
	return s
}

func (s *GetQualityProjectLogResponseBodyData) SetProjectCreateTime(v string) *GetQualityProjectLogResponseBodyData {
	s.ProjectCreateTime = &v
	return s
}

func (s *GetQualityProjectLogResponseBodyData) SetProjectId(v int64) *GetQualityProjectLogResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetQualityProjectLogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
