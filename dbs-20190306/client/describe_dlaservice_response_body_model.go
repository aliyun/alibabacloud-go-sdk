// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDLAServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoAdd(v bool) *DescribeDLAServiceResponseBody
	GetAutoAdd() *bool
	SetErrCode(v string) *DescribeDLAServiceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDLAServiceResponseBody
	GetErrMessage() *string
	SetHaveJobFailed(v bool) *DescribeDLAServiceResponseBody
	GetHaveJobFailed() *bool
	SetHttpStatusCode(v int32) *DescribeDLAServiceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeDLAServiceResponseBody
	GetRequestId() *string
	SetState(v string) *DescribeDLAServiceResponseBody
	GetState() *string
	SetSuccess(v bool) *DescribeDLAServiceResponseBody
	GetSuccess() *bool
}

type DescribeDLAServiceResponseBody struct {
	// Specifies whether to enable the feature of automatically adding incremental data to a data lake. If this feature is enabled, DBS adds the backup sets that are newly generated to the data lake that is created for the backup schedule. Valid values:
	//
	// 	- **true**: enables the feature.
	//
	// 	- **false**: disables the feature.
	//
	// example:
	//
	// true
	AutoAdd *bool `json:"AutoAdd,omitempty" xml:"AutoAdd,omitempty"`
	// The error code.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// Indicates whether a failed DLA task exists in the return result. Valid values:
	//
	// 	- **true**: A failed DLA task exists.
	//
	// 	- **false**: No failed DLA task exists.
	//
	// example:
	//
	// false
	HaveJobFailed *bool `json:"HaveJobFailed,omitempty" xml:"HaveJobFailed,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F1888AC-1138-4995-B9FE-D2734F61C058
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the DLA service for the backup schedule. Valid values:
	//
	// 	- **Running**: DLA is running.
	//
	// 	- **Closing**: DLA is being disabled.
	//
	// 	- **Closed**: DLA is disabled.
	//
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDLAServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDLAServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDLAServiceResponseBody) GetAutoAdd() *bool {
	return s.AutoAdd
}

func (s *DescribeDLAServiceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDLAServiceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDLAServiceResponseBody) GetHaveJobFailed() *bool {
	return s.HaveJobFailed
}

func (s *DescribeDLAServiceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDLAServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDLAServiceResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeDLAServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDLAServiceResponseBody) SetAutoAdd(v bool) *DescribeDLAServiceResponseBody {
	s.AutoAdd = &v
	return s
}

func (s *DescribeDLAServiceResponseBody) SetErrCode(v string) *DescribeDLAServiceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDLAServiceResponseBody) SetErrMessage(v string) *DescribeDLAServiceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDLAServiceResponseBody) SetHaveJobFailed(v bool) *DescribeDLAServiceResponseBody {
	s.HaveJobFailed = &v
	return s
}

func (s *DescribeDLAServiceResponseBody) SetHttpStatusCode(v int32) *DescribeDLAServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDLAServiceResponseBody) SetRequestId(v string) *DescribeDLAServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDLAServiceResponseBody) SetState(v string) *DescribeDLAServiceResponseBody {
	s.State = &v
	return s
}

func (s *DescribeDLAServiceResponseBody) SetSuccess(v bool) *DescribeDLAServiceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDLAServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
