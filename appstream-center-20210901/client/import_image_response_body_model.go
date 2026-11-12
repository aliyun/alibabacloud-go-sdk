// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportImageResponseBody
	GetCode() *string
	SetData(v *ImportImageResponseBodyData) *ImportImageResponseBody
	GetData() *ImportImageResponseBodyData
	SetMessage(v string) *ImportImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportImageResponseBody
	GetSuccess() *bool
}

type ImportImageResponseBody struct {
	// The return code of the API. A value of 200 indicates success. If the call fails, the corresponding error code is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The image identifiers returned for the import request.
	Data *ImportImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The processing message returned by the API.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which is used to locate and troubleshoot the request.
	//
	// example:
	//
	// 8737D130-BFD0-5D51-96F6-C08EB1139A25
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call was successful. Valid values:
	//
	// - true: successful.
	//
	// - false: failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportImageResponseBody) GoString() string {
	return s.String()
}

func (s *ImportImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportImageResponseBody) GetData() *ImportImageResponseBodyData {
	return s.Data
}

func (s *ImportImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportImageResponseBody) SetCode(v string) *ImportImageResponseBody {
	s.Code = &v
	return s
}

func (s *ImportImageResponseBody) SetData(v *ImportImageResponseBodyData) *ImportImageResponseBody {
	s.Data = v
	return s
}

func (s *ImportImageResponseBody) SetMessage(v string) *ImportImageResponseBody {
	s.Message = &v
	return s
}

func (s *ImportImageResponseBody) SetRequestId(v string) *ImportImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportImageResponseBody) SetSuccess(v bool) *ImportImageResponseBody {
	s.Success = &v
	return s
}

func (s *ImportImageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportImageResponseBodyData struct {
	// The image ID.
	//
	// example:
	//
	// imgc-xxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image version ID.
	//
	// example:
	//
	// iv-xxx
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ImportImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImportImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportImageResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *ImportImageResponseBodyData) GetVersionId() *string {
	return s.VersionId
}

func (s *ImportImageResponseBodyData) SetImageId(v string) *ImportImageResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *ImportImageResponseBodyData) SetVersionId(v string) *ImportImageResponseBodyData {
	s.VersionId = &v
	return s
}

func (s *ImportImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
