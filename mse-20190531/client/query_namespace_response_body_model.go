// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryNamespaceResponseBodyData) *QueryNamespaceResponseBody
	GetData() []*QueryNamespaceResponseBodyData
	SetErrorCode(v string) *QueryNamespaceResponseBody
	GetErrorCode() *string
	SetMessage(v string) *QueryNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryNamespaceResponseBody
	GetSuccess() *bool
}

type QueryNamespaceResponseBody struct {
	// The data returned.
	Data []*QueryNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// abcde-fg
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryNamespaceResponseBody) GetData() []*QueryNamespaceResponseBodyData {
	return s.Data
}

func (s *QueryNamespaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryNamespaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryNamespaceResponseBody) SetData(v []*QueryNamespaceResponseBodyData) *QueryNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *QueryNamespaceResponseBody) SetErrorCode(v string) *QueryNamespaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryNamespaceResponseBody) SetMessage(v string) *QueryNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *QueryNamespaceResponseBody) SetRequestId(v string) *QueryNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryNamespaceResponseBody) SetSuccess(v bool) *QueryNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryNamespaceResponseBody) Validate() error {
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

type QueryNamespaceResponseBodyData struct {
	// The name of the namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region to which the namespace belongs.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s QueryNamespaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryNamespaceResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *QueryNamespaceResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *QueryNamespaceResponseBodyData) SetNamespace(v string) *QueryNamespaceResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *QueryNamespaceResponseBodyData) SetRegion(v string) *QueryNamespaceResponseBodyData {
	s.Region = &v
	return s
}

func (s *QueryNamespaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
