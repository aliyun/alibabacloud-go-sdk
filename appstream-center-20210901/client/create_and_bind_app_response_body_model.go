// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndBindAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAndBindAppResponseBody
	GetCode() *string
	SetData(v *CreateAndBindAppResponseBodyData) *CreateAndBindAppResponseBody
	GetData() *CreateAndBindAppResponseBodyData
	SetMessage(v string) *CreateAndBindAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAndBindAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAndBindAppResponseBody
	GetSuccess() *bool
}

type CreateAndBindAppResponseBody struct {
	// The error code. This field is returned when the request fails and may not be included in a successful response.
	//
	// example:
	//
	// InternalError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned by this operation.
	Data *CreateAndBindAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error or result message. This field describes the cause when the request fails and may not be included in a successful response.
	//
	// example:
	//
	// Internal error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request. You can provide this value for troubleshooting.
	//
	// example:
	//
	// 01A1A962-5503-107B-AC8F-EBA85DF905AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was processed successfully. A value of true indicates that the request was processed successfully. To check whether an asynchronous task is completed, query the task status separately.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAndBindAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAndBindAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAndBindAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAndBindAppResponseBody) GetData() *CreateAndBindAppResponseBodyData {
	return s.Data
}

func (s *CreateAndBindAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAndBindAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAndBindAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAndBindAppResponseBody) SetCode(v string) *CreateAndBindAppResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAndBindAppResponseBody) SetData(v *CreateAndBindAppResponseBodyData) *CreateAndBindAppResponseBody {
	s.Data = v
	return s
}

func (s *CreateAndBindAppResponseBody) SetMessage(v string) *CreateAndBindAppResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAndBindAppResponseBody) SetRequestId(v string) *CreateAndBindAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAndBindAppResponseBody) SetSuccess(v bool) *CreateAndBindAppResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAndBindAppResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAndBindAppResponseBodyData struct {
	// The application ID, used for subsequent application queries or to dissociate the application from the image.
	//
	// example:
	//
	// app-bp1example****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application version ID, used to locate a specified version of the application. This field is not the image version ID.
	//
	// example:
	//
	// apv-bp1example****
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	// The image ID, used for subsequent image queries or modifications.
	//
	// example:
	//
	// imgc-bp1example****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image version ID, used to query or manage a specified version of the image.
	//
	// example:
	//
	// iv-bp1example****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateAndBindAppResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAndBindAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAndBindAppResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *CreateAndBindAppResponseBodyData) GetAppVersionId() *string {
	return s.AppVersionId
}

func (s *CreateAndBindAppResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *CreateAndBindAppResponseBodyData) GetVersionId() *string {
	return s.VersionId
}

func (s *CreateAndBindAppResponseBodyData) SetAppId(v string) *CreateAndBindAppResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateAndBindAppResponseBodyData) SetAppVersionId(v string) *CreateAndBindAppResponseBodyData {
	s.AppVersionId = &v
	return s
}

func (s *CreateAndBindAppResponseBodyData) SetImageId(v string) *CreateAndBindAppResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *CreateAndBindAppResponseBodyData) SetVersionId(v string) *CreateAndBindAppResponseBodyData {
	s.VersionId = &v
	return s
}

func (s *CreateAndBindAppResponseBodyData) Validate() error {
	return dara.Validate(s)
}
