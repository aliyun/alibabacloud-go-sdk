// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetUserIdByOpenDingtalkIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchGetUserIdByOpenDingtalkIdResponseBody
	GetRequestId() *string
	SetResults(v []*BatchGetUserIdByOpenDingtalkIdResponseBodyResults) *BatchGetUserIdByOpenDingtalkIdResponseBody
	GetResults() []*BatchGetUserIdByOpenDingtalkIdResponseBodyResults
	SetVendorRequestId(v string) *BatchGetUserIdByOpenDingtalkIdResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *BatchGetUserIdByOpenDingtalkIdResponseBody
	GetVendorType() *string
}

type BatchGetUserIdByOpenDingtalkIdResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// []
	Results []*BatchGetUserIdByOpenDingtalkIdResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s BatchGetUserIdByOpenDingtalkIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetUserIdByOpenDingtalkIdResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetUserIdByOpenDingtalkIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetUserIdByOpenDingtalkIdResponseBody) GetResults() []*BatchGetUserIdByOpenDingtalkIdResponseBodyResults {
	return s.Results
}

func (s *BatchGetUserIdByOpenDingtalkIdResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *BatchGetUserIdByOpenDingtalkIdResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *BatchGetUserIdByOpenDingtalkIdResponseBody) SetRequestId(v string) *BatchGetUserIdByOpenDingtalkIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdResponseBody) SetResults(v []*BatchGetUserIdByOpenDingtalkIdResponseBodyResults) *BatchGetUserIdByOpenDingtalkIdResponseBody {
	s.Results = v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdResponseBody) SetVendorRequestId(v string) *BatchGetUserIdByOpenDingtalkIdResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdResponseBody) SetVendorType(v string) *BatchGetUserIdByOpenDingtalkIdResponseBody {
	s.VendorType = &v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdResponseBody) Validate() error {
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

type BatchGetUserIdByOpenDingtalkIdResponseBodyResults struct {
	// example:
	//
	// User.not.found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// DTOJdYJ2IQC4HuexhtjsS8Qxxxx
	OpenDingtalkId *string `json:"OpenDingtalkId,omitempty" xml:"OpenDingtalkId,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BatchGetUserIdByOpenDingtalkIdResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchGetUserIdByOpenDingtalkIdResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchGetUserIdByOpenDingtalkIdResponseBodyResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BatchGetUserIdByOpenDingtalkIdResponseBodyResults) GetOpenDingtalkId() *string {
	return s.OpenDingtalkId
}

func (s *BatchGetUserIdByOpenDingtalkIdResponseBodyResults) GetUserId() *string {
	return s.UserId
}

func (s *BatchGetUserIdByOpenDingtalkIdResponseBodyResults) SetErrorMessage(v string) *BatchGetUserIdByOpenDingtalkIdResponseBodyResults {
	s.ErrorMessage = &v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdResponseBodyResults) SetOpenDingtalkId(v string) *BatchGetUserIdByOpenDingtalkIdResponseBodyResults {
	s.OpenDingtalkId = &v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdResponseBodyResults) SetUserId(v string) *BatchGetUserIdByOpenDingtalkIdResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
