// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaClusterNamespaceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOpaClusterNamespaceListResponseBody
	GetCode() *string
	SetCount(v int32) *GetOpaClusterNamespaceListResponseBody
	GetCount() *int32
	SetData(v []*GetOpaClusterNamespaceListResponseBodyData) *GetOpaClusterNamespaceListResponseBody
	GetData() []*GetOpaClusterNamespaceListResponseBodyData
	SetMessage(v string) *GetOpaClusterNamespaceListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOpaClusterNamespaceListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOpaClusterNamespaceListResponseBody
	GetSuccess() *bool
}

type GetOpaClusterNamespaceListResponseBody struct {
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
	// 0
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The namespaces.
	Data []*GetOpaClusterNamespaceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 384BFAF1-FC41-58DD-97DD-9D361ADF377D
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

func (s GetOpaClusterNamespaceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOpaClusterNamespaceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpaClusterNamespaceListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOpaClusterNamespaceListResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *GetOpaClusterNamespaceListResponseBody) GetData() []*GetOpaClusterNamespaceListResponseBodyData {
	return s.Data
}

func (s *GetOpaClusterNamespaceListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOpaClusterNamespaceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpaClusterNamespaceListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOpaClusterNamespaceListResponseBody) SetCode(v string) *GetOpaClusterNamespaceListResponseBody {
	s.Code = &v
	return s
}

func (s *GetOpaClusterNamespaceListResponseBody) SetCount(v int32) *GetOpaClusterNamespaceListResponseBody {
	s.Count = &v
	return s
}

func (s *GetOpaClusterNamespaceListResponseBody) SetData(v []*GetOpaClusterNamespaceListResponseBodyData) *GetOpaClusterNamespaceListResponseBody {
	s.Data = v
	return s
}

func (s *GetOpaClusterNamespaceListResponseBody) SetMessage(v string) *GetOpaClusterNamespaceListResponseBody {
	s.Message = &v
	return s
}

func (s *GetOpaClusterNamespaceListResponseBody) SetRequestId(v string) *GetOpaClusterNamespaceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpaClusterNamespaceListResponseBody) SetSuccess(v bool) *GetOpaClusterNamespaceListResponseBody {
	s.Success = &v
	return s
}

func (s *GetOpaClusterNamespaceListResponseBody) Validate() error {
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

type GetOpaClusterNamespaceListResponseBodyData struct {
	// The name of the namespace.
	//
	// example:
	//
	// testNameSpace
	NameSpaceName *string `json:"NameSpaceName,omitempty" xml:"NameSpaceName,omitempty"`
}

func (s GetOpaClusterNamespaceListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOpaClusterNamespaceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOpaClusterNamespaceListResponseBodyData) GetNameSpaceName() *string {
	return s.NameSpaceName
}

func (s *GetOpaClusterNamespaceListResponseBodyData) SetNameSpaceName(v string) *GetOpaClusterNamespaceListResponseBodyData {
	s.NameSpaceName = &v
	return s
}

func (s *GetOpaClusterNamespaceListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
