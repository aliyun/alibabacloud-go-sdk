// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceRecordV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteFaceRecordV2ResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteFaceRecordV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFaceRecordV2ResponseBody
	GetRequestId() *string
	SetResult(v *DeleteFaceRecordV2ResponseBodyResult) *DeleteFaceRecordV2ResponseBody
	GetResult() *DeleteFaceRecordV2ResponseBodyResult
}

type DeleteFaceRecordV2ResponseBody struct {
	// The return code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5E63B760-0ECB-5C07-8503-A65C27876968
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response result.
	Result *DeleteFaceRecordV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteFaceRecordV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceRecordV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceRecordV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFaceRecordV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFaceRecordV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFaceRecordV2ResponseBody) GetResult() *DeleteFaceRecordV2ResponseBodyResult {
	return s.Result
}

func (s *DeleteFaceRecordV2ResponseBody) SetCode(v string) *DeleteFaceRecordV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceRecordV2ResponseBody) SetMessage(v string) *DeleteFaceRecordV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFaceRecordV2ResponseBody) SetRequestId(v string) *DeleteFaceRecordV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceRecordV2ResponseBody) SetResult(v *DeleteFaceRecordV2ResponseBodyResult) *DeleteFaceRecordV2ResponseBody {
	s.Result = v
	return s
}

func (s *DeleteFaceRecordV2ResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteFaceRecordV2ResponseBodyResult struct {
	// The deletion result. Valid values:
	//
	// - Y: Succeeded.
	//
	// - N: Failed.
	//
	// example:
	//
	// Y
	Deleted *string `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// The list of face group codes from which the face data was actually deleted (comma-separated). This parameter is returned with all deleted group codes when FaceGroupCode is not specified.
	//
	// example:
	//
	// wqe***,dsa***
	DeletedGroupCodes *string `json:"DeletedGroupCodes,omitempty" xml:"DeletedGroupCodes,omitempty"`
}

func (s DeleteFaceRecordV2ResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceRecordV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteFaceRecordV2ResponseBodyResult) GetDeleted() *string {
	return s.Deleted
}

func (s *DeleteFaceRecordV2ResponseBodyResult) GetDeletedGroupCodes() *string {
	return s.DeletedGroupCodes
}

func (s *DeleteFaceRecordV2ResponseBodyResult) SetDeleted(v string) *DeleteFaceRecordV2ResponseBodyResult {
	s.Deleted = &v
	return s
}

func (s *DeleteFaceRecordV2ResponseBodyResult) SetDeletedGroupCodes(v string) *DeleteFaceRecordV2ResponseBodyResult {
	s.DeletedGroupCodes = &v
	return s
}

func (s *DeleteFaceRecordV2ResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
