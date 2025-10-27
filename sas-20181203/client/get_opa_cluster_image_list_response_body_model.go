// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaClusterImageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOpaClusterImageListResponseBody
	GetCode() *string
	SetCount(v int32) *GetOpaClusterImageListResponseBody
	GetCount() *int32
	SetData(v []*GetOpaClusterImageListResponseBodyData) *GetOpaClusterImageListResponseBody
	GetData() []*GetOpaClusterImageListResponseBodyData
	SetMessage(v string) *GetOpaClusterImageListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOpaClusterImageListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOpaClusterImageListResponseBody
	GetSuccess() *bool
}

type GetOpaClusterImageListResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of images returned on the current page.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The information about the images.
	Data []*GetOpaClusterImageListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0D02F593-2050-5F5D-8C98-D965FF1B461D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOpaClusterImageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOpaClusterImageListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpaClusterImageListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOpaClusterImageListResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *GetOpaClusterImageListResponseBody) GetData() []*GetOpaClusterImageListResponseBodyData {
	return s.Data
}

func (s *GetOpaClusterImageListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOpaClusterImageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpaClusterImageListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOpaClusterImageListResponseBody) SetCode(v string) *GetOpaClusterImageListResponseBody {
	s.Code = &v
	return s
}

func (s *GetOpaClusterImageListResponseBody) SetCount(v int32) *GetOpaClusterImageListResponseBody {
	s.Count = &v
	return s
}

func (s *GetOpaClusterImageListResponseBody) SetData(v []*GetOpaClusterImageListResponseBodyData) *GetOpaClusterImageListResponseBody {
	s.Data = v
	return s
}

func (s *GetOpaClusterImageListResponseBody) SetMessage(v string) *GetOpaClusterImageListResponseBody {
	s.Message = &v
	return s
}

func (s *GetOpaClusterImageListResponseBody) SetRequestId(v string) *GetOpaClusterImageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpaClusterImageListResponseBody) SetSuccess(v bool) *GetOpaClusterImageListResponseBody {
	s.Success = &v
	return s
}

func (s *GetOpaClusterImageListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOpaClusterImageListResponseBodyData struct {
	// The name of the image.
	//
	// example:
	//
	// opa-test
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
}

func (s GetOpaClusterImageListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOpaClusterImageListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOpaClusterImageListResponseBodyData) GetImageName() *string {
	return s.ImageName
}

func (s *GetOpaClusterImageListResponseBodyData) SetImageName(v string) *GetOpaClusterImageListResponseBodyData {
	s.ImageName = &v
	return s
}

func (s *GetOpaClusterImageListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
