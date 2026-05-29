// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAutomaticWriteOffChangeRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAutomaticWriteOffChangeRecordsResponseBody
	GetCode() *string
	SetData(v []*QueryAutomaticWriteOffChangeRecordsResponseBodyData) *QueryAutomaticWriteOffChangeRecordsResponseBody
	GetData() []*QueryAutomaticWriteOffChangeRecordsResponseBodyData
	SetMessage(v string) *QueryAutomaticWriteOffChangeRecordsResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAutomaticWriteOffChangeRecordsResponseBody
	GetRequestId() *string
}

type QueryAutomaticWriteOffChangeRecordsResponseBody struct {
	// example:
	//
	// 200
	Code    *string                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    []*QueryAutomaticWriteOffChangeRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message *string                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 8cd24f2917797624314748873d0096
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAutomaticWriteOffChangeRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAutomaticWriteOffChangeRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBody) GetData() []*QueryAutomaticWriteOffChangeRecordsResponseBodyData {
	return s.Data
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBody) SetCode(v string) *QueryAutomaticWriteOffChangeRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBody) SetData(v []*QueryAutomaticWriteOffChangeRecordsResponseBodyData) *QueryAutomaticWriteOffChangeRecordsResponseBody {
	s.Data = v
	return s
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBody) SetMessage(v string) *QueryAutomaticWriteOffChangeRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBody) SetRequestId(v string) *QueryAutomaticWriteOffChangeRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBody) Validate() error {
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

type QueryAutomaticWriteOffChangeRecordsResponseBodyData struct {
	ChangeContent *string `json:"ChangeContent,omitempty" xml:"ChangeContent,omitempty"`
	// example:
	//
	// 12312312
	OperateId *string `json:"OperateId,omitempty" xml:"OperateId,omitempty"`
	// example:
	//
	// API
	OperateSource *string `json:"OperateSource,omitempty" xml:"OperateSource,omitempty"`
	// example:
	//
	// 2023-12-15 10:34:36 UTC+8
	OperationTime *string `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
}

func (s QueryAutomaticWriteOffChangeRecordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAutomaticWriteOffChangeRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBodyData) GetChangeContent() *string {
	return s.ChangeContent
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBodyData) GetOperateId() *string {
	return s.OperateId
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBodyData) GetOperateSource() *string {
	return s.OperateSource
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBodyData) GetOperationTime() *string {
	return s.OperationTime
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBodyData) SetChangeContent(v string) *QueryAutomaticWriteOffChangeRecordsResponseBodyData {
	s.ChangeContent = &v
	return s
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBodyData) SetOperateId(v string) *QueryAutomaticWriteOffChangeRecordsResponseBodyData {
	s.OperateId = &v
	return s
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBodyData) SetOperateSource(v string) *QueryAutomaticWriteOffChangeRecordsResponseBodyData {
	s.OperateSource = &v
	return s
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBodyData) SetOperationTime(v string) *QueryAutomaticWriteOffChangeRecordsResponseBodyData {
	s.OperationTime = &v
	return s
}

func (s *QueryAutomaticWriteOffChangeRecordsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
