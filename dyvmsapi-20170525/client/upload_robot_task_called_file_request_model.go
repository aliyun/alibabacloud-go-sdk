// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadRobotTaskCalledFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalledNumber(v string) *UploadRobotTaskCalledFileRequest
	GetCalledNumber() *string
	SetId(v int64) *UploadRobotTaskCalledFileRequest
	GetId() *int64
	SetOwnerId(v int64) *UploadRobotTaskCalledFileRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UploadRobotTaskCalledFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UploadRobotTaskCalledFileRequest
	GetResourceOwnerId() *int64
	SetTtsParam(v string) *UploadRobotTaskCalledFileRequest
	GetTtsParam() *string
	SetTtsParamHead(v string) *UploadRobotTaskCalledFileRequest
	GetTtsParamHead() *string
}

type UploadRobotTaskCalledFileRequest struct {
	// The called numbers. Separate multiple called numbers with commas (,).
	//
	// > After you create a robocall task, you must upload called numbers in batches. You can upload up to 300,000 called numbers for each task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1370****000,1370****111
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The unique ID of the robocall task. You can call the [CreateRobotTask](~~CreateRobotTask~~) operation to obtain the ID of the robocall task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1045****
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The values of the variable in the text-to-speech (TTS) template, in the JSON format. The variable values specified by the TtsParam parameter must match the variable names specified by the TtsParamHead parameter.
	//
	// 	- If all the called numbers carry the same variable values, you can set the value of the number field to **all*	- and upload only one copy of the variable values.
	//
	// 	- If only some of the called numbers carry the same variable values, you can set the value of the number field to **all*	- for these called numbers and set the value of the number field and variable values for other called numbers based on your business requirements. The system preferentially selects the values that you set for the called numbers.
	//
	// example:
	//
	// [{"number":"1370****000","params":["xiaowang","xiaoli","xiaozhou"]}]
	TtsParam *string `json:"TtsParam,omitempty" xml:"TtsParam,omitempty"`
	// The list of variable names carried in the robocall task, in the JSON format. The TtsParamHead parameter must be used together with the TtsParam parameter.
	//
	// example:
	//
	// ["name1","name2","name3"]
	TtsParamHead *string `json:"TtsParamHead,omitempty" xml:"TtsParamHead,omitempty"`
}

func (s UploadRobotTaskCalledFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadRobotTaskCalledFileRequest) GoString() string {
	return s.String()
}

func (s *UploadRobotTaskCalledFileRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *UploadRobotTaskCalledFileRequest) GetId() *int64 {
	return s.Id
}

func (s *UploadRobotTaskCalledFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UploadRobotTaskCalledFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UploadRobotTaskCalledFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UploadRobotTaskCalledFileRequest) GetTtsParam() *string {
	return s.TtsParam
}

func (s *UploadRobotTaskCalledFileRequest) GetTtsParamHead() *string {
	return s.TtsParamHead
}

func (s *UploadRobotTaskCalledFileRequest) SetCalledNumber(v string) *UploadRobotTaskCalledFileRequest {
	s.CalledNumber = &v
	return s
}

func (s *UploadRobotTaskCalledFileRequest) SetId(v int64) *UploadRobotTaskCalledFileRequest {
	s.Id = &v
	return s
}

func (s *UploadRobotTaskCalledFileRequest) SetOwnerId(v int64) *UploadRobotTaskCalledFileRequest {
	s.OwnerId = &v
	return s
}

func (s *UploadRobotTaskCalledFileRequest) SetResourceOwnerAccount(v string) *UploadRobotTaskCalledFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UploadRobotTaskCalledFileRequest) SetResourceOwnerId(v int64) *UploadRobotTaskCalledFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UploadRobotTaskCalledFileRequest) SetTtsParam(v string) *UploadRobotTaskCalledFileRequest {
	s.TtsParam = &v
	return s
}

func (s *UploadRobotTaskCalledFileRequest) SetTtsParamHead(v string) *UploadRobotTaskCalledFileRequest {
	s.TtsParamHead = &v
	return s
}

func (s *UploadRobotTaskCalledFileRequest) Validate() error {
	return dara.Validate(s)
}
