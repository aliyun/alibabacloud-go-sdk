// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbolishCrossProjectPipelineRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AbolishCrossProjectPipelineRunResponseBodyData) *AbolishCrossProjectPipelineRunResponseBody
	GetData() *AbolishCrossProjectPipelineRunResponseBodyData
	SetRequestId(v string) *AbolishCrossProjectPipelineRunResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AbolishCrossProjectPipelineRunResponseBody
	GetSuccess() *bool
}

type AbolishCrossProjectPipelineRunResponseBody struct {
	// The business response.
	//
	// example:
	//
	// {"RequestId":"735894D1-D5E5-50B8-8A6D-041C90A98B23"}
	Data *AbolishCrossProjectPipelineRunResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID, which is used to locate and troubleshoot this API call.
	//
	// example:
	//
	// 735894D1-D5E5-50B8-8A6D-041C90A98B23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AbolishCrossProjectPipelineRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AbolishCrossProjectPipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *AbolishCrossProjectPipelineRunResponseBody) GetData() *AbolishCrossProjectPipelineRunResponseBodyData {
	return s.Data
}

func (s *AbolishCrossProjectPipelineRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AbolishCrossProjectPipelineRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AbolishCrossProjectPipelineRunResponseBody) SetData(v *AbolishCrossProjectPipelineRunResponseBodyData) *AbolishCrossProjectPipelineRunResponseBody {
	s.Data = v
	return s
}

func (s *AbolishCrossProjectPipelineRunResponseBody) SetRequestId(v string) *AbolishCrossProjectPipelineRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbolishCrossProjectPipelineRunResponseBody) SetSuccess(v bool) *AbolishCrossProjectPipelineRunResponseBody {
	s.Success = &v
	return s
}

func (s *AbolishCrossProjectPipelineRunResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AbolishCrossProjectPipelineRunResponseBodyData struct {
	// The request ID.
	//
	// example:
	//
	// 735894D1-D5E5-50B8-8A6D-041C90A98B23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AbolishCrossProjectPipelineRunResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AbolishCrossProjectPipelineRunResponseBodyData) GoString() string {
	return s.String()
}

func (s *AbolishCrossProjectPipelineRunResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *AbolishCrossProjectPipelineRunResponseBodyData) SetRequestId(v string) *AbolishCrossProjectPipelineRunResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *AbolishCrossProjectPipelineRunResponseBodyData) Validate() error {
	return dara.Validate(s)
}
