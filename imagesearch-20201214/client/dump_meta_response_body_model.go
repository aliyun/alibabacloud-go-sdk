// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDumpMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DumpMetaResponseBodyData) *DumpMetaResponseBody
	GetData() *DumpMetaResponseBodyData
	SetRequestId(v string) *DumpMetaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DumpMetaResponseBody
	GetSuccess() *bool
}

type DumpMetaResponseBody struct {
	// The information about the export task.
	Data *DumpMetaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 36C43E96-8F68-44AA-B1AF-B1F7AB94A6C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DumpMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DumpMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DumpMetaResponseBody) GetData() *DumpMetaResponseBodyData {
	return s.Data
}

func (s *DumpMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DumpMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DumpMetaResponseBody) SetData(v *DumpMetaResponseBodyData) *DumpMetaResponseBody {
	s.Data = v
	return s
}

func (s *DumpMetaResponseBody) SetRequestId(v string) *DumpMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DumpMetaResponseBody) SetSuccess(v bool) *DumpMetaResponseBody {
	s.Success = &v
	return s
}

func (s *DumpMetaResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DumpMetaResponseBodyData struct {
	// The status of the export task.
	//
	// 	- PROCESSING: in progress
	//
	// 	- FAIL: failed
	//
	// 	- SUCCESS: successful
	//
	// example:
	//
	// PROCESSING
	DumpMetaStatus *string `json:"DumpMetaStatus,omitempty" xml:"DumpMetaStatus,omitempty"`
	// The ID of the export task.
	//
	// example:
	//
	// 500
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DumpMetaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DumpMetaResponseBodyData) GoString() string {
	return s.String()
}

func (s *DumpMetaResponseBodyData) GetDumpMetaStatus() *string {
	return s.DumpMetaStatus
}

func (s *DumpMetaResponseBodyData) GetId() *string {
	return s.Id
}

func (s *DumpMetaResponseBodyData) SetDumpMetaStatus(v string) *DumpMetaResponseBodyData {
	s.DumpMetaStatus = &v
	return s
}

func (s *DumpMetaResponseBodyData) SetId(v string) *DumpMetaResponseBodyData {
	s.Id = &v
	return s
}

func (s *DumpMetaResponseBodyData) Validate() error {
	return dara.Validate(s)
}
