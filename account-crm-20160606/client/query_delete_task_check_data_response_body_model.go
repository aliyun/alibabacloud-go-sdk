// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeleteTaskCheckDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryDeleteTaskCheckDataResponseBody
	GetCode() *string
	SetMessage(v string) *QueryDeleteTaskCheckDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryDeleteTaskCheckDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryDeleteTaskCheckDataResponseBody
	GetSuccess() *bool
	SetTaskCheckDataDtoList(v []*QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList) *QueryDeleteTaskCheckDataResponseBody
	GetTaskCheckDataDtoList() []*QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList
}

type QueryDeleteTaskCheckDataResponseBody struct {
	Code                 *string                                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message              *string                                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success              *bool                                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskCheckDataDtoList []*QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList `json:"TaskCheckDataDtoList,omitempty" xml:"TaskCheckDataDtoList,omitempty" type:"Repeated"`
}

func (s QueryDeleteTaskCheckDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDeleteTaskCheckDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeleteTaskCheckDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryDeleteTaskCheckDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryDeleteTaskCheckDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDeleteTaskCheckDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDeleteTaskCheckDataResponseBody) GetTaskCheckDataDtoList() []*QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList {
	return s.TaskCheckDataDtoList
}

func (s *QueryDeleteTaskCheckDataResponseBody) SetCode(v string) *QueryDeleteTaskCheckDataResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeleteTaskCheckDataResponseBody) SetMessage(v string) *QueryDeleteTaskCheckDataResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDeleteTaskCheckDataResponseBody) SetRequestId(v string) *QueryDeleteTaskCheckDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeleteTaskCheckDataResponseBody) SetSuccess(v bool) *QueryDeleteTaskCheckDataResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDeleteTaskCheckDataResponseBody) SetTaskCheckDataDtoList(v []*QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList) *QueryDeleteTaskCheckDataResponseBody {
	s.TaskCheckDataDtoList = v
	return s
}

func (s *QueryDeleteTaskCheckDataResponseBody) Validate() error {
	if s.TaskCheckDataDtoList != nil {
		for _, item := range s.TaskCheckDataDtoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList struct {
	CheckerDesc     *string `json:"CheckerDesc,omitempty" xml:"CheckerDesc,omitempty"`
	CheckerName     *string `json:"CheckerName,omitempty" xml:"CheckerName,omitempty"`
	CheckerUniKey   *string `json:"CheckerUniKey,omitempty" xml:"CheckerUniKey,omitempty"`
	DependencyLevel *string `json:"DependencyLevel,omitempty" xml:"DependencyLevel,omitempty"`
}

func (s QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList) String() string {
	return dara.Prettify(s)
}

func (s QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList) GoString() string {
	return s.String()
}

func (s *QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList) GetCheckerDesc() *string {
	return s.CheckerDesc
}

func (s *QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList) GetCheckerName() *string {
	return s.CheckerName
}

func (s *QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList) GetCheckerUniKey() *string {
	return s.CheckerUniKey
}

func (s *QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList) GetDependencyLevel() *string {
	return s.DependencyLevel
}

func (s *QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList) SetCheckerDesc(v string) *QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList {
	s.CheckerDesc = &v
	return s
}

func (s *QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList) SetCheckerName(v string) *QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList {
	s.CheckerName = &v
	return s
}

func (s *QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList) SetCheckerUniKey(v string) *QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList {
	s.CheckerUniKey = &v
	return s
}

func (s *QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList) SetDependencyLevel(v string) *QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList {
	s.DependencyLevel = &v
	return s
}

func (s *QueryDeleteTaskCheckDataResponseBodyTaskCheckDataDtoList) Validate() error {
	return dara.Validate(s)
}
