// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUploadSkillsViaOssResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *BatchUploadSkillsViaOssResponseBodyData) *BatchUploadSkillsViaOssResponseBody
	GetData() *BatchUploadSkillsViaOssResponseBodyData
	SetRequestId(v string) *BatchUploadSkillsViaOssResponseBody
	GetRequestId() *string
}

type BatchUploadSkillsViaOssResponseBody struct {
	// The response data.
	Data *BatchUploadSkillsViaOssResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchUploadSkillsViaOssResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUploadSkillsViaOssResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUploadSkillsViaOssResponseBody) GetData() *BatchUploadSkillsViaOssResponseBodyData {
	return s.Data
}

func (s *BatchUploadSkillsViaOssResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUploadSkillsViaOssResponseBody) SetData(v *BatchUploadSkillsViaOssResponseBodyData) *BatchUploadSkillsViaOssResponseBody {
	s.Data = v
	return s
}

func (s *BatchUploadSkillsViaOssResponseBody) SetRequestId(v string) *BatchUploadSkillsViaOssResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUploadSkillsViaOssResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchUploadSkillsViaOssResponseBodyData struct {
	// The batch upload results.
	Results []*BatchUploadSkillsViaOssResponseBodyDataResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s BatchUploadSkillsViaOssResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchUploadSkillsViaOssResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchUploadSkillsViaOssResponseBodyData) GetResults() []*BatchUploadSkillsViaOssResponseBodyDataResults {
	return s.Results
}

func (s *BatchUploadSkillsViaOssResponseBodyData) SetResults(v []*BatchUploadSkillsViaOssResponseBodyDataResults) *BatchUploadSkillsViaOssResponseBodyData {
	s.Results = v
	return s
}

func (s *BatchUploadSkillsViaOssResponseBodyData) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchUploadSkillsViaOssResponseBodyDataResults struct {
	// The error code.
	//
	// example:
	//
	// VALIDATION_FAILED
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter validation failed
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The name.
	//
	// example:
	//
	// skill-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The resource owner.
	//
	// example:
	//
	// alice
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchUploadSkillsViaOssResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s BatchUploadSkillsViaOssResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *BatchUploadSkillsViaOssResponseBodyDataResults) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BatchUploadSkillsViaOssResponseBodyDataResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BatchUploadSkillsViaOssResponseBodyDataResults) GetName() *string {
	return s.Name
}

func (s *BatchUploadSkillsViaOssResponseBodyDataResults) GetOwner() *string {
	return s.Owner
}

func (s *BatchUploadSkillsViaOssResponseBodyDataResults) GetSuccess() *bool {
	return s.Success
}

func (s *BatchUploadSkillsViaOssResponseBodyDataResults) SetErrorCode(v string) *BatchUploadSkillsViaOssResponseBodyDataResults {
	s.ErrorCode = &v
	return s
}

func (s *BatchUploadSkillsViaOssResponseBodyDataResults) SetErrorMessage(v string) *BatchUploadSkillsViaOssResponseBodyDataResults {
	s.ErrorMessage = &v
	return s
}

func (s *BatchUploadSkillsViaOssResponseBodyDataResults) SetName(v string) *BatchUploadSkillsViaOssResponseBodyDataResults {
	s.Name = &v
	return s
}

func (s *BatchUploadSkillsViaOssResponseBodyDataResults) SetOwner(v string) *BatchUploadSkillsViaOssResponseBodyDataResults {
	s.Owner = &v
	return s
}

func (s *BatchUploadSkillsViaOssResponseBodyDataResults) SetSuccess(v bool) *BatchUploadSkillsViaOssResponseBodyDataResults {
	s.Success = &v
	return s
}

func (s *BatchUploadSkillsViaOssResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}
