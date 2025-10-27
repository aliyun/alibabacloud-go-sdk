// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCanTrySasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetCanTrySasResponseBodyData) *GetCanTrySasResponseBody
	GetData() *GetCanTrySasResponseBodyData
	SetRequestId(v string) *GetCanTrySasResponseBody
	GetRequestId() *string
}

type GetCanTrySasResponseBody struct {
	// The data returned.
	Data *GetCanTrySasResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8BAA57***B7073A5C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCanTrySasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCanTrySasResponseBody) GoString() string {
	return s.String()
}

func (s *GetCanTrySasResponseBody) GetData() *GetCanTrySasResponseBodyData {
	return s.Data
}

func (s *GetCanTrySasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCanTrySasResponseBody) SetData(v *GetCanTrySasResponseBodyData) *GetCanTrySasResponseBody {
	s.Data = v
	return s
}

func (s *GetCanTrySasResponseBody) SetRequestId(v string) *GetCanTrySasResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCanTrySasResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCanTrySasResponseBodyData struct {
	// Indicates whether the user is qualified for the trial use. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	CanTry *int32 `json:"CanTry,omitempty" xml:"CanTry,omitempty"`
	// The editions that are allowed for the trial use.
	CanTryVersions []*int32 `json:"CanTryVersions,omitempty" xml:"CanTryVersions,omitempty" type:"Repeated"`
	// The trial type. Valid values:
	//
	// 	- **0**: trial prohibited
	//
	// 	- **1**: first trial
	//
	// 	- **2**: second trial
	//
	// example:
	//
	// 1
	TryType *int32 `json:"TryType,omitempty" xml:"TryType,omitempty"`
}

func (s GetCanTrySasResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCanTrySasResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCanTrySasResponseBodyData) GetCanTry() *int32 {
	return s.CanTry
}

func (s *GetCanTrySasResponseBodyData) GetCanTryVersions() []*int32 {
	return s.CanTryVersions
}

func (s *GetCanTrySasResponseBodyData) GetTryType() *int32 {
	return s.TryType
}

func (s *GetCanTrySasResponseBodyData) SetCanTry(v int32) *GetCanTrySasResponseBodyData {
	s.CanTry = &v
	return s
}

func (s *GetCanTrySasResponseBodyData) SetCanTryVersions(v []*int32) *GetCanTrySasResponseBodyData {
	s.CanTryVersions = v
	return s
}

func (s *GetCanTrySasResponseBodyData) SetTryType(v int32) *GetCanTrySasResponseBodyData {
	s.TryType = &v
	return s
}

func (s *GetCanTrySasResponseBodyData) Validate() error {
	return dara.Validate(s)
}
