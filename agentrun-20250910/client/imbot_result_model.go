// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIMBotResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IMBotResult
	GetCode() *string
	SetData(v *IMBotInfo) *IMBotResult
	GetData() *IMBotInfo
	SetRequestId(v string) *IMBotResult
	GetRequestId() *string
}

type IMBotResult struct {
	// OK 表示成功
	Code *string    `json:"code,omitempty" xml:"code,omitempty"`
	Data *IMBotInfo `json:"data,omitempty" xml:"data,omitempty"`
	// 与响应头 x-funagent-request-id 对应
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s IMBotResult) String() string {
	return dara.Prettify(s)
}

func (s IMBotResult) GoString() string {
	return s.String()
}

func (s *IMBotResult) GetCode() *string {
	return s.Code
}

func (s *IMBotResult) GetData() *IMBotInfo {
	return s.Data
}

func (s *IMBotResult) GetRequestId() *string {
	return s.RequestId
}

func (s *IMBotResult) SetCode(v string) *IMBotResult {
	s.Code = &v
	return s
}

func (s *IMBotResult) SetData(v *IMBotInfo) *IMBotResult {
	s.Data = v
	return s
}

func (s *IMBotResult) SetRequestId(v string) *IMBotResult {
	s.RequestId = &v
	return s
}

func (s *IMBotResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
