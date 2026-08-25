// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *PublishImageResponseBodyData) *PublishImageResponseBody
	GetData() *PublishImageResponseBodyData
	SetRequestId(v string) *PublishImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PublishImageResponseBody
	GetSuccess() *bool
}

type PublishImageResponseBody struct {
	// The result of the API request.
	//
	// example:
	//
	// true
	Data *PublishImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PublishImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishImageResponseBody) GoString() string {
	return s.String()
}

func (s *PublishImageResponseBody) GetData() *PublishImageResponseBodyData {
	return s.Data
}

func (s *PublishImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PublishImageResponseBody) SetData(v *PublishImageResponseBodyData) *PublishImageResponseBody {
	s.Data = v
	return s
}

func (s *PublishImageResponseBody) SetRequestId(v string) *PublishImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishImageResponseBody) SetSuccess(v bool) *PublishImageResponseBody {
	s.Success = &v
	return s
}

func (s *PublishImageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PublishImageResponseBodyData struct {
	// The image publish execution ID.
	//
	// example:
	//
	// 582d4896-d224-413b-b883-239eeebe0bc5
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// Indicates whether the trigger was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PublishImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PublishImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *PublishImageResponseBodyData) GetProcessId() *string {
	return s.ProcessId
}

func (s *PublishImageResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *PublishImageResponseBodyData) SetProcessId(v string) *PublishImageResponseBodyData {
	s.ProcessId = &v
	return s
}

func (s *PublishImageResponseBodyData) SetSuccess(v bool) *PublishImageResponseBodyData {
	s.Success = &v
	return s
}

func (s *PublishImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
