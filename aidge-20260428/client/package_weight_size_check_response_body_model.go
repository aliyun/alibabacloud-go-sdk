// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPackageWeightSizeCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PackageWeightSizeCheckResponseBody
	GetCode() *string
	SetData(v *PackageWeightSizeCheckResponseBodyData) *PackageWeightSizeCheckResponseBody
	GetData() *PackageWeightSizeCheckResponseBodyData
	SetMessage(v string) *PackageWeightSizeCheckResponseBody
	GetMessage() *string
	SetRequestId(v string) *PackageWeightSizeCheckResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PackageWeightSizeCheckResponseBody
	GetSuccess() *bool
}

type PackageWeightSizeCheckResponseBody struct {
	// The response code. 200 indicates a successful call. For other response codes, refer to the error code information.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The asynchronous submission result data, which contains the asynchronous task ID.
	Data *PackageWeightSizeCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. "Success" is returned for normal calls, and specific error information is returned for exceptions.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, used to identify a unique request call.
	//
	// example:
	//
	// F93D82E4-D0B6-1043-AC58-282597BC3C10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. true indicates success, and false indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PackageWeightSizeCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PackageWeightSizeCheckResponseBody) GoString() string {
	return s.String()
}

func (s *PackageWeightSizeCheckResponseBody) GetCode() *string {
	return s.Code
}

func (s *PackageWeightSizeCheckResponseBody) GetData() *PackageWeightSizeCheckResponseBodyData {
	return s.Data
}

func (s *PackageWeightSizeCheckResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PackageWeightSizeCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PackageWeightSizeCheckResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PackageWeightSizeCheckResponseBody) SetCode(v string) *PackageWeightSizeCheckResponseBody {
	s.Code = &v
	return s
}

func (s *PackageWeightSizeCheckResponseBody) SetData(v *PackageWeightSizeCheckResponseBodyData) *PackageWeightSizeCheckResponseBody {
	s.Data = v
	return s
}

func (s *PackageWeightSizeCheckResponseBody) SetMessage(v string) *PackageWeightSizeCheckResponseBody {
	s.Message = &v
	return s
}

func (s *PackageWeightSizeCheckResponseBody) SetRequestId(v string) *PackageWeightSizeCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *PackageWeightSizeCheckResponseBody) SetSuccess(v bool) *PackageWeightSizeCheckResponseBody {
	s.Success = &v
	return s
}

func (s *PackageWeightSizeCheckResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PackageWeightSizeCheckResponseBodyData struct {
	// The asynchronous task ID, used to query the audit result later through QueryAsyncTaskResult.
	//
	// example:
	//
	// a8323ada-a196-9061-976f-90e38b27323a
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PackageWeightSizeCheckResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PackageWeightSizeCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *PackageWeightSizeCheckResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *PackageWeightSizeCheckResponseBodyData) SetTaskId(v string) *PackageWeightSizeCheckResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *PackageWeightSizeCheckResponseBodyData) Validate() error {
	return dara.Validate(s)
}
