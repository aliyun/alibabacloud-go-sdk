// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaClusterLabelListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOpaClusterLabelListResponseBody
	GetCode() *string
	SetCount(v int32) *GetOpaClusterLabelListResponseBody
	GetCount() *int32
	SetData(v []*GetOpaClusterLabelListResponseBodyData) *GetOpaClusterLabelListResponseBody
	GetData() []*GetOpaClusterLabelListResponseBodyData
	SetMessage(v string) *GetOpaClusterLabelListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOpaClusterLabelListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOpaClusterLabelListResponseBody
	GetSuccess() *bool
}

type GetOpaClusterLabelListResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The information about the tags that are added to containers.
	Data []*GetOpaClusterLabelListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 95D35EB3-1F8E-5E07-A68E-BE018C9B80CB
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

func (s GetOpaClusterLabelListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOpaClusterLabelListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpaClusterLabelListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOpaClusterLabelListResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *GetOpaClusterLabelListResponseBody) GetData() []*GetOpaClusterLabelListResponseBodyData {
	return s.Data
}

func (s *GetOpaClusterLabelListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOpaClusterLabelListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpaClusterLabelListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOpaClusterLabelListResponseBody) SetCode(v string) *GetOpaClusterLabelListResponseBody {
	s.Code = &v
	return s
}

func (s *GetOpaClusterLabelListResponseBody) SetCount(v int32) *GetOpaClusterLabelListResponseBody {
	s.Count = &v
	return s
}

func (s *GetOpaClusterLabelListResponseBody) SetData(v []*GetOpaClusterLabelListResponseBodyData) *GetOpaClusterLabelListResponseBody {
	s.Data = v
	return s
}

func (s *GetOpaClusterLabelListResponseBody) SetMessage(v string) *GetOpaClusterLabelListResponseBody {
	s.Message = &v
	return s
}

func (s *GetOpaClusterLabelListResponseBody) SetRequestId(v string) *GetOpaClusterLabelListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpaClusterLabelListResponseBody) SetSuccess(v bool) *GetOpaClusterLabelListResponseBody {
	s.Success = &v
	return s
}

func (s *GetOpaClusterLabelListResponseBody) Validate() error {
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

type GetOpaClusterLabelListResponseBodyData struct {
	// The name of the tag that is added to the container.
	//
	// example:
	//
	// test
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s GetOpaClusterLabelListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOpaClusterLabelListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOpaClusterLabelListResponseBodyData) GetTagName() *string {
	return s.TagName
}

func (s *GetOpaClusterLabelListResponseBodyData) SetTagName(v string) *GetOpaClusterLabelListResponseBodyData {
	s.TagName = &v
	return s
}

func (s *GetOpaClusterLabelListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
