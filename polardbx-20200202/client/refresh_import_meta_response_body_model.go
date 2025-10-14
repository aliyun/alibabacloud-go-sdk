// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshImportMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RefreshImportMetaResponseBodyData) *RefreshImportMetaResponseBody
	GetData() *RefreshImportMetaResponseBodyData
	SetMessage(v string) *RefreshImportMetaResponseBody
	GetMessage() *string
	SetRequestId(v string) *RefreshImportMetaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RefreshImportMetaResponseBody
	GetSuccess() *bool
}

type RefreshImportMetaResponseBody struct {
	Data *RefreshImportMetaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AE4F6C34-065F-45AA-B5DC-4B8D816F6305
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RefreshImportMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshImportMetaResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshImportMetaResponseBody) GetData() *RefreshImportMetaResponseBodyData {
	return s.Data
}

func (s *RefreshImportMetaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RefreshImportMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshImportMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RefreshImportMetaResponseBody) SetData(v *RefreshImportMetaResponseBodyData) *RefreshImportMetaResponseBody {
	s.Data = v
	return s
}

func (s *RefreshImportMetaResponseBody) SetMessage(v string) *RefreshImportMetaResponseBody {
	s.Message = &v
	return s
}

func (s *RefreshImportMetaResponseBody) SetRequestId(v string) *RefreshImportMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshImportMetaResponseBody) SetSuccess(v bool) *RefreshImportMetaResponseBody {
	s.Success = &v
	return s
}

func (s *RefreshImportMetaResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RefreshImportMetaResponseBodyData struct {
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
}

func (s RefreshImportMetaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RefreshImportMetaResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefreshImportMetaResponseBodyData) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *RefreshImportMetaResponseBodyData) SetSlinkTaskId(v string) *RefreshImportMetaResponseBodyData {
	s.SlinkTaskId = &v
	return s
}

func (s *RefreshImportMetaResponseBodyData) Validate() error {
	return dara.Validate(s)
}
