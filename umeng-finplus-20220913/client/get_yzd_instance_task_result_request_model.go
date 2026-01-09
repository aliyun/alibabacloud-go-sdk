// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYzdInstanceTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *GetYzdInstanceTaskResultRequestBody) *GetYzdInstanceTaskResultRequest
	GetBody() *GetYzdInstanceTaskResultRequestBody
}

type GetYzdInstanceTaskResultRequest struct {
	Body *GetYzdInstanceTaskResultRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s GetYzdInstanceTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetYzdInstanceTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetYzdInstanceTaskResultRequest) GetBody() *GetYzdInstanceTaskResultRequestBody {
	return s.Body
}

func (s *GetYzdInstanceTaskResultRequest) SetBody(v *GetYzdInstanceTaskResultRequestBody) *GetYzdInstanceTaskResultRequest {
	s.Body = v
	return s
}

func (s *GetYzdInstanceTaskResultRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetYzdInstanceTaskResultRequestBody struct {
	AppId              *string `json:"appId,omitempty" xml:"appId,omitempty"`
	RangeTimeEndTime   *string `json:"rangeTimeEndTime,omitempty" xml:"rangeTimeEndTime,omitempty"`
	RangeTimeStartTime *string `json:"rangeTimeStartTime,omitempty" xml:"rangeTimeStartTime,omitempty"`
}

func (s GetYzdInstanceTaskResultRequestBody) String() string {
	return dara.Prettify(s)
}

func (s GetYzdInstanceTaskResultRequestBody) GoString() string {
	return s.String()
}

func (s *GetYzdInstanceTaskResultRequestBody) GetAppId() *string {
	return s.AppId
}

func (s *GetYzdInstanceTaskResultRequestBody) GetRangeTimeEndTime() *string {
	return s.RangeTimeEndTime
}

func (s *GetYzdInstanceTaskResultRequestBody) GetRangeTimeStartTime() *string {
	return s.RangeTimeStartTime
}

func (s *GetYzdInstanceTaskResultRequestBody) SetAppId(v string) *GetYzdInstanceTaskResultRequestBody {
	s.AppId = &v
	return s
}

func (s *GetYzdInstanceTaskResultRequestBody) SetRangeTimeEndTime(v string) *GetYzdInstanceTaskResultRequestBody {
	s.RangeTimeEndTime = &v
	return s
}

func (s *GetYzdInstanceTaskResultRequestBody) SetRangeTimeStartTime(v string) *GetYzdInstanceTaskResultRequestBody {
	s.RangeTimeStartTime = &v
	return s
}

func (s *GetYzdInstanceTaskResultRequestBody) Validate() error {
	return dara.Validate(s)
}
