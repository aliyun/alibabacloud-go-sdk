// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectHttpApiConflictsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DetectHttpApiConflictsResponseBody
	GetCode() *string
	SetData(v *DetectHttpApiConflictsResponseBodyData) *DetectHttpApiConflictsResponseBody
	GetData() *DetectHttpApiConflictsResponseBodyData
	SetMessage(v string) *DetectHttpApiConflictsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DetectHttpApiConflictsResponseBody
	GetRequestId() *string
}

type DetectHttpApiConflictsResponseBody struct {
	Code      *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data      *DetectHttpApiConflictsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DetectHttpApiConflictsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectHttpApiConflictsResponseBody) GoString() string {
	return s.String()
}

func (s *DetectHttpApiConflictsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DetectHttpApiConflictsResponseBody) GetData() *DetectHttpApiConflictsResponseBodyData {
	return s.Data
}

func (s *DetectHttpApiConflictsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetectHttpApiConflictsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectHttpApiConflictsResponseBody) SetCode(v string) *DetectHttpApiConflictsResponseBody {
	s.Code = &v
	return s
}

func (s *DetectHttpApiConflictsResponseBody) SetData(v *DetectHttpApiConflictsResponseBodyData) *DetectHttpApiConflictsResponseBody {
	s.Data = v
	return s
}

func (s *DetectHttpApiConflictsResponseBody) SetMessage(v string) *DetectHttpApiConflictsResponseBody {
	s.Message = &v
	return s
}

func (s *DetectHttpApiConflictsResponseBody) SetRequestId(v string) *DetectHttpApiConflictsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectHttpApiConflictsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetectHttpApiConflictsResponseBodyData struct {
	Items []*ApiRouteConflictInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s DetectHttpApiConflictsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectHttpApiConflictsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectHttpApiConflictsResponseBodyData) GetItems() []*ApiRouteConflictInfo {
	return s.Items
}

func (s *DetectHttpApiConflictsResponseBodyData) SetItems(v []*ApiRouteConflictInfo) *DetectHttpApiConflictsResponseBodyData {
	s.Items = v
	return s
}

func (s *DetectHttpApiConflictsResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
