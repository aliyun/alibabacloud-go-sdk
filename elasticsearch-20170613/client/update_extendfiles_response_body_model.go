// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExtendfilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateExtendfilesResponseBody
	GetRequestId() *string
	SetResult(v []*UpdateExtendfilesResponseBodyResult) *UpdateExtendfilesResponseBody
	GetResult() []*UpdateExtendfilesResponseBodyResult
}

type UpdateExtendfilesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*UpdateExtendfilesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s UpdateExtendfilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExtendfilesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExtendfilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExtendfilesResponseBody) GetResult() []*UpdateExtendfilesResponseBodyResult {
	return s.Result
}

func (s *UpdateExtendfilesResponseBody) SetRequestId(v string) *UpdateExtendfilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExtendfilesResponseBody) SetResult(v []*UpdateExtendfilesResponseBodyResult) *UpdateExtendfilesResponseBody {
	s.Result = v
	return s
}

func (s *UpdateExtendfilesResponseBody) Validate() error {
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

type UpdateExtendfilesResponseBodyResult struct {
	// The size of the driver file. Unit: byte.
	//
	// example:
	//
	// 1853083
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// The name of the driver file.
	//
	// example:
	//
	// mysql-connector-java-6.0.2.jar
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The source of the driver file. This parameter is fixed as ORIGIN, which indicates that the driver file is retained.
	//
	// example:
	//
	// ORIGIN
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s UpdateExtendfilesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateExtendfilesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateExtendfilesResponseBodyResult) GetFileSize() *int64 {
	return s.FileSize
}

func (s *UpdateExtendfilesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateExtendfilesResponseBodyResult) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateExtendfilesResponseBodyResult) SetFileSize(v int64) *UpdateExtendfilesResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *UpdateExtendfilesResponseBodyResult) SetName(v string) *UpdateExtendfilesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateExtendfilesResponseBodyResult) SetSourceType(v string) *UpdateExtendfilesResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *UpdateExtendfilesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
