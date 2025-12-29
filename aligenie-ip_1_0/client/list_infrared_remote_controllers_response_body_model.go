// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInfraredRemoteControllersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ListInfraredRemoteControllersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInfraredRemoteControllersResponseBody
	GetRequestId() *string
	SetResult(v []*ListInfraredRemoteControllersResponseBodyResult) *ListInfraredRemoteControllersResponseBody
	GetResult() []*ListInfraredRemoteControllersResponseBodyResult
	SetStatusCode(v int32) *ListInfraredRemoteControllersResponseBody
	GetStatusCode() *int32
}

type ListInfraredRemoteControllersResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0C90A059-3653-5356-A78E-8A6BDA606155
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListInfraredRemoteControllersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListInfraredRemoteControllersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInfraredRemoteControllersResponseBody) GoString() string {
	return s.String()
}

func (s *ListInfraredRemoteControllersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInfraredRemoteControllersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInfraredRemoteControllersResponseBody) GetResult() []*ListInfraredRemoteControllersResponseBodyResult {
	return s.Result
}

func (s *ListInfraredRemoteControllersResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInfraredRemoteControllersResponseBody) SetMessage(v string) *ListInfraredRemoteControllersResponseBody {
	s.Message = &v
	return s
}

func (s *ListInfraredRemoteControllersResponseBody) SetRequestId(v string) *ListInfraredRemoteControllersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInfraredRemoteControllersResponseBody) SetResult(v []*ListInfraredRemoteControllersResponseBodyResult) *ListInfraredRemoteControllersResponseBody {
	s.Result = v
	return s
}

func (s *ListInfraredRemoteControllersResponseBody) SetStatusCode(v int32) *ListInfraredRemoteControllersResponseBody {
	s.StatusCode = &v
	return s
}

func (s *ListInfraredRemoteControllersResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInfraredRemoteControllersResponseBodyResult struct {
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// 3747
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// 4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListInfraredRemoteControllersResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListInfraredRemoteControllersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInfraredRemoteControllersResponseBodyResult) GetIndex() *int32 {
	return s.Index
}

func (s *ListInfraredRemoteControllersResponseBodyResult) GetRid() *int64 {
	return s.Rid
}

func (s *ListInfraredRemoteControllersResponseBodyResult) GetVersion() *string {
	return s.Version
}

func (s *ListInfraredRemoteControllersResponseBodyResult) SetIndex(v int32) *ListInfraredRemoteControllersResponseBodyResult {
	s.Index = &v
	return s
}

func (s *ListInfraredRemoteControllersResponseBodyResult) SetRid(v int64) *ListInfraredRemoteControllersResponseBodyResult {
	s.Rid = &v
	return s
}

func (s *ListInfraredRemoteControllersResponseBodyResult) SetVersion(v string) *ListInfraredRemoteControllersResponseBodyResult {
	s.Version = &v
	return s
}

func (s *ListInfraredRemoteControllersResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
