// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseExpiredPrepaidInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClass(v string) *ReleaseExpiredPrepaidInstanceResponseBody
	GetClass() *string
	SetCode(v string) *ReleaseExpiredPrepaidInstanceResponseBody
	GetCode() *string
	SetData(v *ReleaseExpiredPrepaidInstanceResponseBodyData) *ReleaseExpiredPrepaidInstanceResponseBody
	GetData() *ReleaseExpiredPrepaidInstanceResponseBodyData
	SetMessage(v string) *ReleaseExpiredPrepaidInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReleaseExpiredPrepaidInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReleaseExpiredPrepaidInstanceResponseBody
	GetSuccess() *bool
}

type ReleaseExpiredPrepaidInstanceResponseBody struct {
	// **[Deprecated]*	- The field is deprecated.
	//
	// example:
	//
	// deprecated
	Class *string `json:"Class,omitempty" xml:"Class,omitempty"`
	// The error code returned when the call fails. For more information, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ReleaseExpiredPrepaidInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DF58589C-A06C-4224-8615-7797E6474FA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseExpiredPrepaidInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseExpiredPrepaidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseExpiredPrepaidInstanceResponseBody) GetClass() *string {
	return s.Class
}

func (s *ReleaseExpiredPrepaidInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReleaseExpiredPrepaidInstanceResponseBody) GetData() *ReleaseExpiredPrepaidInstanceResponseBodyData {
	return s.Data
}

func (s *ReleaseExpiredPrepaidInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReleaseExpiredPrepaidInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseExpiredPrepaidInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReleaseExpiredPrepaidInstanceResponseBody) SetClass(v string) *ReleaseExpiredPrepaidInstanceResponseBody {
	s.Class = &v
	return s
}

func (s *ReleaseExpiredPrepaidInstanceResponseBody) SetCode(v string) *ReleaseExpiredPrepaidInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseExpiredPrepaidInstanceResponseBody) SetData(v *ReleaseExpiredPrepaidInstanceResponseBodyData) *ReleaseExpiredPrepaidInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ReleaseExpiredPrepaidInstanceResponseBody) SetMessage(v string) *ReleaseExpiredPrepaidInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseExpiredPrepaidInstanceResponseBody) SetRequestId(v string) *ReleaseExpiredPrepaidInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseExpiredPrepaidInstanceResponseBody) SetSuccess(v bool) *ReleaseExpiredPrepaidInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ReleaseExpiredPrepaidInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReleaseExpiredPrepaidInstanceResponseBodyData struct {
	// **[Deprecated]*	- The field is deprecated.
	//
	// example:
	//
	// deprecated
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The list of instance results.
	InstanceResults []*ReleaseExpiredPrepaidInstanceResponseBodyDataInstanceResults `json:"InstanceResults,omitempty" xml:"InstanceResults,omitempty" type:"Repeated"`
	// Indicates whether the release operation is successful.
	//
	// example:
	//
	// true
	ReleaseResult *bool `json:"ReleaseResult,omitempty" xml:"ReleaseResult,omitempty"`
}

func (s ReleaseExpiredPrepaidInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReleaseExpiredPrepaidInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReleaseExpiredPrepaidInstanceResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *ReleaseExpiredPrepaidInstanceResponseBodyData) GetInstanceResults() []*ReleaseExpiredPrepaidInstanceResponseBodyDataInstanceResults {
	return s.InstanceResults
}

func (s *ReleaseExpiredPrepaidInstanceResponseBodyData) GetReleaseResult() *bool {
	return s.ReleaseResult
}

func (s *ReleaseExpiredPrepaidInstanceResponseBodyData) SetHostId(v string) *ReleaseExpiredPrepaidInstanceResponseBodyData {
	s.HostId = &v
	return s
}

func (s *ReleaseExpiredPrepaidInstanceResponseBodyData) SetInstanceResults(v []*ReleaseExpiredPrepaidInstanceResponseBodyDataInstanceResults) *ReleaseExpiredPrepaidInstanceResponseBodyData {
	s.InstanceResults = v
	return s
}

func (s *ReleaseExpiredPrepaidInstanceResponseBodyData) SetReleaseResult(v bool) *ReleaseExpiredPrepaidInstanceResponseBodyData {
	s.ReleaseResult = &v
	return s
}

func (s *ReleaseExpiredPrepaidInstanceResponseBodyData) Validate() error {
	if s.InstanceResults != nil {
		for _, item := range s.InstanceResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReleaseExpiredPrepaidInstanceResponseBodyDataInstanceResults struct {
	// The instance ID.
	//
	// example:
	//
	// rm-1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseExpiredPrepaidInstanceResponseBodyDataInstanceResults) String() string {
	return dara.Prettify(s)
}

func (s ReleaseExpiredPrepaidInstanceResponseBodyDataInstanceResults) GoString() string {
	return s.String()
}

func (s *ReleaseExpiredPrepaidInstanceResponseBodyDataInstanceResults) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseExpiredPrepaidInstanceResponseBodyDataInstanceResults) GetSuccess() *bool {
	return s.Success
}

func (s *ReleaseExpiredPrepaidInstanceResponseBodyDataInstanceResults) SetInstanceId(v string) *ReleaseExpiredPrepaidInstanceResponseBodyDataInstanceResults {
	s.InstanceId = &v
	return s
}

func (s *ReleaseExpiredPrepaidInstanceResponseBodyDataInstanceResults) SetSuccess(v bool) *ReleaseExpiredPrepaidInstanceResponseBodyDataInstanceResults {
	s.Success = &v
	return s
}

func (s *ReleaseExpiredPrepaidInstanceResponseBodyDataInstanceResults) Validate() error {
	return dara.Validate(s)
}
