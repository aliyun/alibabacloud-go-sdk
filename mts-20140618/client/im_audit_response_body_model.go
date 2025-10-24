// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImAuditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageQuotaExceed(v bool) *ImAuditResponseBody
	GetImageQuotaExceed() *bool
	SetImageResults(v *ImAuditResponseBodyImageResults) *ImAuditResponseBody
	GetImageResults() *ImAuditResponseBodyImageResults
	SetRequestId(v string) *ImAuditResponseBody
	GetRequestId() *string
	SetTextQuotaExceed(v bool) *ImAuditResponseBody
	GetTextQuotaExceed() *bool
	SetTextResults(v *ImAuditResponseBodyTextResults) *ImAuditResponseBody
	GetTextResults() *ImAuditResponseBodyTextResults
}

type ImAuditResponseBody struct {
	// Indicates whether the image moderation QPS exceeds the limit. Valid values: true and false. A value of true indicates that the QPS does not exceed the limit. A value of false indicates that the QPS exceeds the limit.
	//
	// example:
	//
	// false
	ImageQuotaExceed *bool `json:"ImageQuotaExceed,omitempty" xml:"ImageQuotaExceed,omitempty"`
	// The image moderation results. If the HTTP status code 200 is returned, the array in the returned results contains one or more elements. For more information about the parameters, see [Data returned by the ImAudit operation](https://help.aliyun.com/document_detail/268644.html).
	ImageResults *ImAuditResponseBodyImageResults `json:"ImageResults,omitempty" xml:"ImageResults,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 5210DBB0-E327-4D45-ADBC-0B83C8793421
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the text moderation QPS exceeds the limit. Valid values: true and false.
	//
	// example:
	//
	// false
	TextQuotaExceed *bool `json:"TextQuotaExceed,omitempty" xml:"TextQuotaExceed,omitempty"`
	// The text moderation results. If the HTTP status code 200 is returned, the array in the returned results contains one or more elements. For more information about the parameters, see [Data returned by the ImAudit operation](https://help.aliyun.com/document_detail/268644.html).
	TextResults *ImAuditResponseBodyTextResults `json:"TextResults,omitempty" xml:"TextResults,omitempty" type:"Struct"`
}

func (s ImAuditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImAuditResponseBody) GoString() string {
	return s.String()
}

func (s *ImAuditResponseBody) GetImageQuotaExceed() *bool {
	return s.ImageQuotaExceed
}

func (s *ImAuditResponseBody) GetImageResults() *ImAuditResponseBodyImageResults {
	return s.ImageResults
}

func (s *ImAuditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImAuditResponseBody) GetTextQuotaExceed() *bool {
	return s.TextQuotaExceed
}

func (s *ImAuditResponseBody) GetTextResults() *ImAuditResponseBodyTextResults {
	return s.TextResults
}

func (s *ImAuditResponseBody) SetImageQuotaExceed(v bool) *ImAuditResponseBody {
	s.ImageQuotaExceed = &v
	return s
}

func (s *ImAuditResponseBody) SetImageResults(v *ImAuditResponseBodyImageResults) *ImAuditResponseBody {
	s.ImageResults = v
	return s
}

func (s *ImAuditResponseBody) SetRequestId(v string) *ImAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImAuditResponseBody) SetTextQuotaExceed(v bool) *ImAuditResponseBody {
	s.TextQuotaExceed = &v
	return s
}

func (s *ImAuditResponseBody) SetTextResults(v *ImAuditResponseBodyTextResults) *ImAuditResponseBody {
	s.TextResults = v
	return s
}

func (s *ImAuditResponseBody) Validate() error {
	if s.ImageResults != nil {
		if err := s.ImageResults.Validate(); err != nil {
			return err
		}
	}
	if s.TextResults != nil {
		if err := s.TextResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImAuditResponseBodyImageResults struct {
	// The image moderation results.
	Result []*ImAuditResponseBodyImageResultsResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ImAuditResponseBodyImageResults) String() string {
	return dara.Prettify(s)
}

func (s ImAuditResponseBodyImageResults) GoString() string {
	return s.String()
}

func (s *ImAuditResponseBodyImageResults) GetResult() []*ImAuditResponseBodyImageResultsResult {
	return s.Result
}

func (s *ImAuditResponseBodyImageResults) SetResult(v []*ImAuditResponseBodyImageResultsResult) *ImAuditResponseBodyImageResults {
	s.Result = v
	return s
}

func (s *ImAuditResponseBodyImageResults) Validate() error {
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

type ImAuditResponseBodyImageResultsResult struct {
	// The error code. The error code is the same as the HTTP status code. This parameter is not returned if the request is successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// The ID of the moderated object.
	//
	// >  If you set the dataId parameter in the moderation request, the dataId parameter is returned in the response.
	//
	// example:
	//
	// uuid-1234-1234-1234
	DataId *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// The additional information about the image. If ad is specified for the Scenes parameter, the following content may be returned for this parameter: hitLibInfo: the information about the custom text library that is hit by the text in the image. The value of this parameter is an array. For more information about the structure, see [hitLibInfo](https://help.aliyun.com/document_detail/268644.html).
	Extras map[string]interface{} `json:"extras,omitempty" xml:"extras,omitempty"`
	// The message that is returned for the request.
	//
	// example:
	//
	// ok
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// The returned data. If the call is successful, the array in the returned results contains one or more elements. Each element is a struct.
	Results []*ImAuditResponseBodyImageResultsResultResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// The ID of the moderation task.
	//
	// example:
	//
	// img4wlJcb7p4wH4lAP3111111-12****
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The URL of the moderated object.
	//
	// example:
	//
	// http://example.com/example-****.jpg
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ImAuditResponseBodyImageResultsResult) String() string {
	return dara.Prettify(s)
}

func (s ImAuditResponseBodyImageResultsResult) GoString() string {
	return s.String()
}

func (s *ImAuditResponseBodyImageResultsResult) GetCode() *int64 {
	return s.Code
}

func (s *ImAuditResponseBodyImageResultsResult) GetDataId() *string {
	return s.DataId
}

func (s *ImAuditResponseBodyImageResultsResult) GetExtras() map[string]interface{} {
	return s.Extras
}

func (s *ImAuditResponseBodyImageResultsResult) GetMsg() *string {
	return s.Msg
}

func (s *ImAuditResponseBodyImageResultsResult) GetResults() []*ImAuditResponseBodyImageResultsResultResults {
	return s.Results
}

func (s *ImAuditResponseBodyImageResultsResult) GetTaskId() *string {
	return s.TaskId
}

func (s *ImAuditResponseBodyImageResultsResult) GetUrl() *string {
	return s.Url
}

func (s *ImAuditResponseBodyImageResultsResult) SetCode(v int64) *ImAuditResponseBodyImageResultsResult {
	s.Code = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResult) SetDataId(v string) *ImAuditResponseBodyImageResultsResult {
	s.DataId = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResult) SetExtras(v map[string]interface{}) *ImAuditResponseBodyImageResultsResult {
	s.Extras = v
	return s
}

func (s *ImAuditResponseBodyImageResultsResult) SetMsg(v string) *ImAuditResponseBodyImageResultsResult {
	s.Msg = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResult) SetResults(v []*ImAuditResponseBodyImageResultsResultResults) *ImAuditResponseBodyImageResultsResult {
	s.Results = v
	return s
}

func (s *ImAuditResponseBodyImageResultsResult) SetTaskId(v string) *ImAuditResponseBodyImageResultsResult {
	s.TaskId = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResult) SetUrl(v string) *ImAuditResponseBodyImageResultsResult {
	s.Url = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResult) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImAuditResponseBodyImageResultsResultResults struct {
	// The category of the moderation results. Valid values vary based on the specified moderation scenario.
	//
	// 	- If the Scenes parameter is set to porn, the valid values are:
	//
	//     	- normal: no pornographic content
	//
	//     	- sexy: sexy content
	//
	//     	- porn: pornographic content
	//
	// 	- If the Scenes parameter is set to terrorism, the valid values are:
	//
	//     	- normal: no pornographic content
	//
	//     	- bloody: bloody content
	//
	//     	- explosion: explosions and smoke
	//
	//     	- outfit: special costume
	//
	//     	- logo: special logo
	//
	//     	- weapon: weapon
	//
	//     	- politics: political content
	//
	//     	- violence: violence
	//
	//     	- crowd: crowd
	//
	//     	- parade: parade
	//
	//     	- carcrash: car accident
	//
	//     	- flag: flag
	//
	//     	- location: landmark
	//
	//     	- others: other content
	//
	// 	- If the Scenes parameter is set to ad, the valid values are:
	//
	//     	- normal: no pornographic content
	//
	//     	- ad: ad violation
	//
	//     	- politics: politically sensitive content in text
	//
	//     	- porn: pornographic content in text
	//
	//     	- abuse: abuse in text
	//
	//     	- terrorism: terrorist content in text
	//
	//     	- contraband: prohibited content in text
	//
	//     	- spam: junk content in text
	//
	//     	- npx: illegal ad
	//
	//     	- qrcode: QR code
	//
	//     	- programCode: mini program code
	//
	// 	- If the Scenes parameter is set to qrcode, the valid values are:
	//
	//     	- normal: no pornographic content
	//
	//     	- qrcode: QR code
	//
	//     	- programCode: mini program code
	//
	// 	- If the Scenes parameter is set to live, the valid values are:
	//
	//     	- normal: no pornographic content
	//
	//     	- meaningless: no content in the image, such as black or white screen
	//
	//     	- PIP: picture-in-picture
	//
	//     	- smoking: smoking
	//
	//     	- drivelive: live broadcasting in a running vehicle
	//
	// 	- If the Scenes parameter is set to logo, the valid values are:
	//
	//     	- normal: no pornographic content
	//
	//     	- TV: controlled logo
	//
	//     	- trademark: trademark
	//
	// example:
	//
	// sexy
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score of the confidence level. Valid values: 0 to 100. A greater value indicates a higher confidence level. If a value of pass is returned for the suggestion parameter, a higher confidence level indicates a higher probability that the content is normal. If a value of review or block is returned for the suggestion parameter, a higher confidence level indicates a higher probability that the content contains violations.
	//
	// >  This score is for reference only. We strongly recommend that you do not use this score in your business. We recommend that you use the values that are returned for the suggestion, label, and sublabel parameters to determine whether the content contains violations. The sublabel parameter is returned by some operations.
	//
	// example:
	//
	// 91.54
	Rate *float64 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The image moderation scenario. Valid values:
	//
	// 	- porn: pornography
	//
	// 	- terrorism: terrorist content
	//
	// 	- ad: ad violation
	//
	// 	- qrcode: QR code
	//
	// 	- live: undesirable scene
	//
	// 	- logo: special logo
	//
	// example:
	//
	// porn
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The recommended subsequent operation. Valid values:
	//
	// 	- pass: The content passes the moderation. No further actions are required.
	//
	// 	- review: The moderation object contains suspected violations and requires human review.
	//
	// 	- block: The moderation object contains violations. We recommend that you delete or block the object.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// If the temporary access URL of the image is too long, a truncated temporary access URL is returned for each frame.
	Frames []*ImAuditResponseBodyImageResultsResultResultsFrames `json:"frames,omitempty" xml:"frames,omitempty" type:"Repeated"`
	// The information about the term hit by the ad or violation text detected in the moderated image.
	HintWordsInfo []*ImAuditResponseBodyImageResultsResultResultsHintWordsInfo `json:"hintWordsInfo,omitempty" xml:"hintWordsInfo,omitempty" type:"Repeated"`
	// The information about the logo detected in the moderated image.
	LogoData []*ImAuditResponseBodyImageResultsResultResultsLogoData `json:"logoData,omitempty" xml:"logoData,omitempty" type:"Repeated"`
	// ocrData
	OcrData []*string `json:"ocrData,omitempty" xml:"ocrData,omitempty" type:"Repeated"`
	// The location information of the mini program code detected in the moderated image.
	ProgramCodeData []*ImAuditResponseBodyImageResultsResultResultsProgramCodeData `json:"programCodeData,omitempty" xml:"programCodeData,omitempty" type:"Repeated"`
	// The information about the text that is included in the QR code detected in the moderated image.
	QrcodeData []*string `json:"qrcodeData,omitempty" xml:"qrcodeData,omitempty" type:"Repeated"`
	// The coordinates of the QR code detected in the image.
	QrcodeLocations []*ImAuditResponseBodyImageResultsResultResultsQrcodeLocations `json:"qrcodeLocations,omitempty" xml:"qrcodeLocations,omitempty" type:"Repeated"`
	// The information about the terrorist content detected in the moderated image.
	SfaceData []*ImAuditResponseBodyImageResultsResultResultsSfaceData `json:"sfaceData,omitempty" xml:"sfaceData,omitempty" type:"Repeated"`
}

func (s ImAuditResponseBodyImageResultsResultResults) String() string {
	return dara.Prettify(s)
}

func (s ImAuditResponseBodyImageResultsResultResults) GoString() string {
	return s.String()
}

func (s *ImAuditResponseBodyImageResultsResultResults) GetLabel() *string {
	return s.Label
}

func (s *ImAuditResponseBodyImageResultsResultResults) GetRate() *float64 {
	return s.Rate
}

func (s *ImAuditResponseBodyImageResultsResultResults) GetScene() *string {
	return s.Scene
}

func (s *ImAuditResponseBodyImageResultsResultResults) GetSuggestion() *string {
	return s.Suggestion
}

func (s *ImAuditResponseBodyImageResultsResultResults) GetFrames() []*ImAuditResponseBodyImageResultsResultResultsFrames {
	return s.Frames
}

func (s *ImAuditResponseBodyImageResultsResultResults) GetHintWordsInfo() []*ImAuditResponseBodyImageResultsResultResultsHintWordsInfo {
	return s.HintWordsInfo
}

func (s *ImAuditResponseBodyImageResultsResultResults) GetLogoData() []*ImAuditResponseBodyImageResultsResultResultsLogoData {
	return s.LogoData
}

func (s *ImAuditResponseBodyImageResultsResultResults) GetOcrData() []*string {
	return s.OcrData
}

func (s *ImAuditResponseBodyImageResultsResultResults) GetProgramCodeData() []*ImAuditResponseBodyImageResultsResultResultsProgramCodeData {
	return s.ProgramCodeData
}

func (s *ImAuditResponseBodyImageResultsResultResults) GetQrcodeData() []*string {
	return s.QrcodeData
}

func (s *ImAuditResponseBodyImageResultsResultResults) GetQrcodeLocations() []*ImAuditResponseBodyImageResultsResultResultsQrcodeLocations {
	return s.QrcodeLocations
}

func (s *ImAuditResponseBodyImageResultsResultResults) GetSfaceData() []*ImAuditResponseBodyImageResultsResultResultsSfaceData {
	return s.SfaceData
}

func (s *ImAuditResponseBodyImageResultsResultResults) SetLabel(v string) *ImAuditResponseBodyImageResultsResultResults {
	s.Label = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResults) SetRate(v float64) *ImAuditResponseBodyImageResultsResultResults {
	s.Rate = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResults) SetScene(v string) *ImAuditResponseBodyImageResultsResultResults {
	s.Scene = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResults) SetSuggestion(v string) *ImAuditResponseBodyImageResultsResultResults {
	s.Suggestion = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResults) SetFrames(v []*ImAuditResponseBodyImageResultsResultResultsFrames) *ImAuditResponseBodyImageResultsResultResults {
	s.Frames = v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResults) SetHintWordsInfo(v []*ImAuditResponseBodyImageResultsResultResultsHintWordsInfo) *ImAuditResponseBodyImageResultsResultResults {
	s.HintWordsInfo = v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResults) SetLogoData(v []*ImAuditResponseBodyImageResultsResultResultsLogoData) *ImAuditResponseBodyImageResultsResultResults {
	s.LogoData = v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResults) SetOcrData(v []*string) *ImAuditResponseBodyImageResultsResultResults {
	s.OcrData = v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResults) SetProgramCodeData(v []*ImAuditResponseBodyImageResultsResultResultsProgramCodeData) *ImAuditResponseBodyImageResultsResultResults {
	s.ProgramCodeData = v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResults) SetQrcodeData(v []*string) *ImAuditResponseBodyImageResultsResultResults {
	s.QrcodeData = v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResults) SetQrcodeLocations(v []*ImAuditResponseBodyImageResultsResultResultsQrcodeLocations) *ImAuditResponseBodyImageResultsResultResults {
	s.QrcodeLocations = v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResults) SetSfaceData(v []*ImAuditResponseBodyImageResultsResultResultsSfaceData) *ImAuditResponseBodyImageResultsResultResults {
	s.SfaceData = v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResults) Validate() error {
	if s.Frames != nil {
		for _, item := range s.Frames {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HintWordsInfo != nil {
		for _, item := range s.HintWordsInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LogoData != nil {
		for _, item := range s.LogoData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ProgramCodeData != nil {
		for _, item := range s.ProgramCodeData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QrcodeLocations != nil {
		for _, item := range s.QrcodeLocations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SfaceData != nil {
		for _, item := range s.SfaceData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImAuditResponseBodyImageResultsResultResultsFrames struct {
	// The score of the confidence level. Valid values: 0 to 100. A higher confidence level indicates higher reliability of the moderation result. We recommend that you do not use this score in your business.
	//
	// example:
	//
	// 89.85
	Rate *float32 `json:"rate,omitempty" xml:"rate,omitempty"`
	// The temporary access URL of the truncated frame. The URL is valid for 5 minutes.
	//
	// example:
	//
	// http://example.com/test-01.jpg
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ImAuditResponseBodyImageResultsResultResultsFrames) String() string {
	return dara.Prettify(s)
}

func (s ImAuditResponseBodyImageResultsResultResultsFrames) GoString() string {
	return s.String()
}

func (s *ImAuditResponseBodyImageResultsResultResultsFrames) GetRate() *float32 {
	return s.Rate
}

func (s *ImAuditResponseBodyImageResultsResultResultsFrames) GetUrl() *string {
	return s.Url
}

func (s *ImAuditResponseBodyImageResultsResultResultsFrames) SetRate(v float32) *ImAuditResponseBodyImageResultsResultResultsFrames {
	s.Rate = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsFrames) SetUrl(v string) *ImAuditResponseBodyImageResultsResultResultsFrames {
	s.Url = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsFrames) Validate() error {
	return dara.Validate(s)
}

type ImAuditResponseBodyImageResultsResultResultsHintWordsInfo struct {
	// The term hit by the detected text.
	//
	// example:
	//
	// Sensitive words
	Context *string `json:"context,omitempty" xml:"context,omitempty"`
}

func (s ImAuditResponseBodyImageResultsResultResultsHintWordsInfo) String() string {
	return dara.Prettify(s)
}

func (s ImAuditResponseBodyImageResultsResultResultsHintWordsInfo) GoString() string {
	return s.String()
}

func (s *ImAuditResponseBodyImageResultsResultResultsHintWordsInfo) GetContext() *string {
	return s.Context
}

func (s *ImAuditResponseBodyImageResultsResultResultsHintWordsInfo) SetContext(v string) *ImAuditResponseBodyImageResultsResultResultsHintWordsInfo {
	s.Context = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsHintWordsInfo) Validate() error {
	return dara.Validate(s)
}

type ImAuditResponseBodyImageResultsResultResultsLogoData struct {
	// The height of the logo area. Unit: pixel.
	//
	// example:
	//
	// 106
	H *float32 `json:"h,omitempty" xml:"h,omitempty"`
	// The name of the detected logo.
	//
	// example:
	//
	// Hunan TV
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the detected logo. For example, a value of TV indicates a controlled media logo.
	//
	// example:
	//
	// TV
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The width of the logo area. Unit: pixel.
	//
	// example:
	//
	// 106
	W *float32 `json:"w,omitempty" xml:"w,omitempty"`
	// The distance between the upper-left corner of the logo area and the y-axis, with the upper-left corner of the image being the coordinate origin. Unit: pixel.
	//
	// example:
	//
	// 140
	X *float32 `json:"x,omitempty" xml:"x,omitempty"`
	// The distance between the upper-left corner of the logo area and the x-axis, with the upper-left corner of the image being the coordinate origin. Unit: pixel.
	//
	// example:
	//
	// 68
	Y *float32 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s ImAuditResponseBodyImageResultsResultResultsLogoData) String() string {
	return dara.Prettify(s)
}

func (s ImAuditResponseBodyImageResultsResultResultsLogoData) GoString() string {
	return s.String()
}

func (s *ImAuditResponseBodyImageResultsResultResultsLogoData) GetH() *float32 {
	return s.H
}

func (s *ImAuditResponseBodyImageResultsResultResultsLogoData) GetName() *string {
	return s.Name
}

func (s *ImAuditResponseBodyImageResultsResultResultsLogoData) GetType() *string {
	return s.Type
}

func (s *ImAuditResponseBodyImageResultsResultResultsLogoData) GetW() *float32 {
	return s.W
}

func (s *ImAuditResponseBodyImageResultsResultResultsLogoData) GetX() *float32 {
	return s.X
}

func (s *ImAuditResponseBodyImageResultsResultResultsLogoData) GetY() *float32 {
	return s.Y
}

func (s *ImAuditResponseBodyImageResultsResultResultsLogoData) SetH(v float32) *ImAuditResponseBodyImageResultsResultResultsLogoData {
	s.H = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsLogoData) SetName(v string) *ImAuditResponseBodyImageResultsResultResultsLogoData {
	s.Name = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsLogoData) SetType(v string) *ImAuditResponseBodyImageResultsResultResultsLogoData {
	s.Type = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsLogoData) SetW(v float32) *ImAuditResponseBodyImageResultsResultResultsLogoData {
	s.W = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsLogoData) SetX(v float32) *ImAuditResponseBodyImageResultsResultResultsLogoData {
	s.X = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsLogoData) SetY(v float32) *ImAuditResponseBodyImageResultsResultResultsLogoData {
	s.Y = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsLogoData) Validate() error {
	return dara.Validate(s)
}

type ImAuditResponseBodyImageResultsResultResultsProgramCodeData struct {
	// The height of the mini program code area. Unit: pixel.
	//
	// example:
	//
	// 413.0
	H *float32 `json:"h,omitempty" xml:"h,omitempty"`
	// The width of the mini program code area. Unit: pixel.
	//
	// example:
	//
	// 402.0
	W *float32 `json:"w,omitempty" xml:"w,omitempty"`
	// The distance between the upper-left corner of the mini program code area and the y-axis, with the upper-left corner of the image being the coordinate origin. Unit: pixel.
	//
	// example:
	//
	// 11.0
	X *float32 `json:"x,omitempty" xml:"x,omitempty"`
	// The distance between the upper-left corner of the mini program code area and the x-axis, with the upper-left corner of the image being the coordinate origin. Unit: pixel.
	//
	// example:
	//
	// 0.0
	Y *float32 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s ImAuditResponseBodyImageResultsResultResultsProgramCodeData) String() string {
	return dara.Prettify(s)
}

func (s ImAuditResponseBodyImageResultsResultResultsProgramCodeData) GoString() string {
	return s.String()
}

func (s *ImAuditResponseBodyImageResultsResultResultsProgramCodeData) GetH() *float32 {
	return s.H
}

func (s *ImAuditResponseBodyImageResultsResultResultsProgramCodeData) GetW() *float32 {
	return s.W
}

func (s *ImAuditResponseBodyImageResultsResultResultsProgramCodeData) GetX() *float32 {
	return s.X
}

func (s *ImAuditResponseBodyImageResultsResultResultsProgramCodeData) GetY() *float32 {
	return s.Y
}

func (s *ImAuditResponseBodyImageResultsResultResultsProgramCodeData) SetH(v float32) *ImAuditResponseBodyImageResultsResultResultsProgramCodeData {
	s.H = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsProgramCodeData) SetW(v float32) *ImAuditResponseBodyImageResultsResultResultsProgramCodeData {
	s.W = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsProgramCodeData) SetX(v float32) *ImAuditResponseBodyImageResultsResultResultsProgramCodeData {
	s.X = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsProgramCodeData) SetY(v float32) *ImAuditResponseBodyImageResultsResultResultsProgramCodeData {
	s.Y = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsProgramCodeData) Validate() error {
	return dara.Validate(s)
}

type ImAuditResponseBodyImageResultsResultResultsQrcodeLocations struct {
	// The height of the QR code area. Unit: pixel.
	//
	// example:
	//
	// 413.0
	H *float32 `json:"h,omitempty" xml:"h,omitempty"`
	// The URL that the detected QR code points to.
	//
	// example:
	//
	// http://xxx
	Qrcode *string `json:"qrcode,omitempty" xml:"qrcode,omitempty"`
	// The width of the QR code area. Unit: pixel.
	//
	// example:
	//
	// 402.0
	W *float32 `json:"w,omitempty" xml:"w,omitempty"`
	// The distance between the upper-left corner of the QR code area and the y-axis, with the upper-left corner of the image being the coordinate origin. Unit: pixel.
	//
	// example:
	//
	// 11
	X *float32 `json:"x,omitempty" xml:"x,omitempty"`
	// The distance between the upper-left corner of the QR code area and the x-axis, with the upper-left corner of the image being the coordinate origin. Unit: pixel.
	//
	// example:
	//
	// 0
	Y *float32 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s ImAuditResponseBodyImageResultsResultResultsQrcodeLocations) String() string {
	return dara.Prettify(s)
}

func (s ImAuditResponseBodyImageResultsResultResultsQrcodeLocations) GoString() string {
	return s.String()
}

func (s *ImAuditResponseBodyImageResultsResultResultsQrcodeLocations) GetH() *float32 {
	return s.H
}

func (s *ImAuditResponseBodyImageResultsResultResultsQrcodeLocations) GetQrcode() *string {
	return s.Qrcode
}

func (s *ImAuditResponseBodyImageResultsResultResultsQrcodeLocations) GetW() *float32 {
	return s.W
}

func (s *ImAuditResponseBodyImageResultsResultResultsQrcodeLocations) GetX() *float32 {
	return s.X
}

func (s *ImAuditResponseBodyImageResultsResultResultsQrcodeLocations) GetY() *float32 {
	return s.Y
}

func (s *ImAuditResponseBodyImageResultsResultResultsQrcodeLocations) SetH(v float32) *ImAuditResponseBodyImageResultsResultResultsQrcodeLocations {
	s.H = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsQrcodeLocations) SetQrcode(v string) *ImAuditResponseBodyImageResultsResultResultsQrcodeLocations {
	s.Qrcode = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsQrcodeLocations) SetW(v float32) *ImAuditResponseBodyImageResultsResultResultsQrcodeLocations {
	s.W = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsQrcodeLocations) SetX(v float32) *ImAuditResponseBodyImageResultsResultResultsQrcodeLocations {
	s.X = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsQrcodeLocations) SetY(v float32) *ImAuditResponseBodyImageResultsResultResultsQrcodeLocations {
	s.Y = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsQrcodeLocations) Validate() error {
	return dara.Validate(s)
}

type ImAuditResponseBodyImageResultsResultResultsSfaceData struct {
	// The information about the face detected in the moderated image.
	Faces []*ImAuditResponseBodyImageResultsResultResultsSfaceDataFaces `json:"faces,omitempty" xml:"faces,omitempty" type:"Repeated"`
	// The height of the face area. Unit: pixel.
	//
	// example:
	//
	// 121
	H *float32 `json:"h,omitempty" xml:"h,omitempty"`
	// The width of the face area. Unit: pixel.
	//
	// example:
	//
	// 47
	W *float32 `json:"w,omitempty" xml:"w,omitempty"`
	// The distance between the upper-left corner of the face area and the y-axis, with the upper-left corner of the image being the coordinate origin. Unit: pixel.
	//
	// example:
	//
	// 49
	X *float32 `json:"x,omitempty" xml:"x,omitempty"`
	// The distance between the upper-left corner of the face area and the y-axis, with the upper-left corner of the image being the coordinate origin. Unit: pixel.
	//
	// example:
	//
	// 39
	Y *float32 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s ImAuditResponseBodyImageResultsResultResultsSfaceData) String() string {
	return dara.Prettify(s)
}

func (s ImAuditResponseBodyImageResultsResultResultsSfaceData) GoString() string {
	return s.String()
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceData) GetFaces() []*ImAuditResponseBodyImageResultsResultResultsSfaceDataFaces {
	return s.Faces
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceData) GetH() *float32 {
	return s.H
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceData) GetW() *float32 {
	return s.W
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceData) GetX() *float32 {
	return s.X
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceData) GetY() *float32 {
	return s.Y
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceData) SetFaces(v []*ImAuditResponseBodyImageResultsResultResultsSfaceDataFaces) *ImAuditResponseBodyImageResultsResultResultsSfaceData {
	s.Faces = v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceData) SetH(v float32) *ImAuditResponseBodyImageResultsResultResultsSfaceData {
	s.H = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceData) SetW(v float32) *ImAuditResponseBodyImageResultsResultResultsSfaceData {
	s.W = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceData) SetX(v float32) *ImAuditResponseBodyImageResultsResultResultsSfaceData {
	s.X = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceData) SetY(v float32) *ImAuditResponseBodyImageResultsResultResultsSfaceData {
	s.Y = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceData) Validate() error {
	if s.Faces != nil {
		for _, item := range s.Faces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImAuditResponseBodyImageResultsResultResultsSfaceDataFaces struct {
	// The ID of the detected face. The value is a string.
	//
	// example:
	//
	// AliFace_0001234
	Idid *string `json:"idid,omitempty" xml:"idid,omitempty"`
	// This value is a string, which indicates the name of a similar person.
	//
	// example:
	//
	// Name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The score of the confidence level. The value is a float point number. Valid values: 0 to 100. A greater value indicates a higher confidence level for facial recognition.
	//
	// example:
	//
	// 91.54
	Re *float32 `json:"re,omitempty" xml:"re,omitempty"`
}

func (s ImAuditResponseBodyImageResultsResultResultsSfaceDataFaces) String() string {
	return dara.Prettify(s)
}

func (s ImAuditResponseBodyImageResultsResultResultsSfaceDataFaces) GoString() string {
	return s.String()
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceDataFaces) GetIdid() *string {
	return s.Idid
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceDataFaces) GetName() *string {
	return s.Name
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceDataFaces) GetRe() *float32 {
	return s.Re
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceDataFaces) SetIdid(v string) *ImAuditResponseBodyImageResultsResultResultsSfaceDataFaces {
	s.Idid = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceDataFaces) SetName(v string) *ImAuditResponseBodyImageResultsResultResultsSfaceDataFaces {
	s.Name = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceDataFaces) SetRe(v float32) *ImAuditResponseBodyImageResultsResultResultsSfaceDataFaces {
	s.Re = &v
	return s
}

func (s *ImAuditResponseBodyImageResultsResultResultsSfaceDataFaces) Validate() error {
	return dara.Validate(s)
}

type ImAuditResponseBodyTextResults struct {
	// The text moderation results.
	Result []*ImAuditResponseBodyTextResultsResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ImAuditResponseBodyTextResults) String() string {
	return dara.Prettify(s)
}

func (s ImAuditResponseBodyTextResults) GoString() string {
	return s.String()
}

func (s *ImAuditResponseBodyTextResults) GetResult() []*ImAuditResponseBodyTextResultsResult {
	return s.Result
}

func (s *ImAuditResponseBodyTextResults) SetResult(v []*ImAuditResponseBodyTextResultsResult) *ImAuditResponseBodyTextResults {
	s.Result = v
	return s
}

func (s *ImAuditResponseBodyTextResults) Validate() error {
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

type ImAuditResponseBodyTextResultsResult struct {
	// The error code. The error code is the same as the HTTP status code. For more information, see [Error codes](https://help.aliyun.com/document_detail/29254.html).
	//
	// example:
	//
	// 200
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// The text that you specify in the moderation request.
	//
	// example:
	//
	// This is test text.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The sequence number of the text.
	//
	// example:
	//
	// cfd33235-71a4-468b-8137-a5ffe323****
	DataId *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// The message that is returned for the request.
	//
	// example:
	//
	// OK
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// The returned data. If the HTTP status code 200 is returned, the array in the returned results contains one or more elements. Each element is a struct.
	Results []*ImAuditResponseBodyTextResultsResultResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// The ID of the moderation task.
	//
	// example:
	//
	// txt6HB8NQoEbU@5fosnj2xVEM-1t****
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ImAuditResponseBodyTextResultsResult) String() string {
	return dara.Prettify(s)
}

func (s ImAuditResponseBodyTextResultsResult) GoString() string {
	return s.String()
}

func (s *ImAuditResponseBodyTextResultsResult) GetCode() *int64 {
	return s.Code
}

func (s *ImAuditResponseBodyTextResultsResult) GetContent() *string {
	return s.Content
}

func (s *ImAuditResponseBodyTextResultsResult) GetDataId() *string {
	return s.DataId
}

func (s *ImAuditResponseBodyTextResultsResult) GetMsg() *string {
	return s.Msg
}

func (s *ImAuditResponseBodyTextResultsResult) GetResults() []*ImAuditResponseBodyTextResultsResultResults {
	return s.Results
}

func (s *ImAuditResponseBodyTextResultsResult) GetTaskId() *string {
	return s.TaskId
}

func (s *ImAuditResponseBodyTextResultsResult) SetCode(v int64) *ImAuditResponseBodyTextResultsResult {
	s.Code = &v
	return s
}

func (s *ImAuditResponseBodyTextResultsResult) SetContent(v string) *ImAuditResponseBodyTextResultsResult {
	s.Content = &v
	return s
}

func (s *ImAuditResponseBodyTextResultsResult) SetDataId(v string) *ImAuditResponseBodyTextResultsResult {
	s.DataId = &v
	return s
}

func (s *ImAuditResponseBodyTextResultsResult) SetMsg(v string) *ImAuditResponseBodyTextResultsResult {
	s.Msg = &v
	return s
}

func (s *ImAuditResponseBodyTextResultsResult) SetResults(v []*ImAuditResponseBodyTextResultsResultResults) *ImAuditResponseBodyTextResultsResult {
	s.Results = v
	return s
}

func (s *ImAuditResponseBodyTextResultsResult) SetTaskId(v string) *ImAuditResponseBodyTextResultsResult {
	s.TaskId = &v
	return s
}

func (s *ImAuditResponseBodyTextResultsResult) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImAuditResponseBodyTextResultsResultResults struct {
	// The risky content that the moderated text hits. A text entry can hit multiple pieces of risky content.
	Details []*ImAuditResponseBodyTextResultsResultResultsDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	// The category of the moderation result for the moderated text. Valid values:
	//
	// 	- normal: normal content
	//
	// 	- spam: spam
	//
	// 	- ad: ad
	//
	// 	- politics: political content
	//
	// 	- terrorism: terrorist content
	//
	// 	- abuse: abuse
	//
	// 	- porn: pornographic content
	//
	// 	- flood: excessive junk content
	//
	// 	- contraband: prohibited content
	//
	// 	- meaningless: meaningless content
	//
	// 	- customized: custom content, such as a custom keyword
	//
	// example:
	//
	// porn
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// The score of the confidence level. Valid values: 0 to 100. A greater value indicates a higher confidence level. If a value of pass is returned for the suggestion parameter, a higher confidence level indicates a higher probability that the content is normal. If a value of review or block is returned for the suggestion parameter, a higher confidence level indicates a higher probability that the content contains violations.
	//
	// >  This score is for reference only. We strongly recommend that you do not use this score in your business. We recommend that you use the values that are returned for the suggestion, label, and sublabel parameters to determine whether the content contains violations. The sublabel parameter is returned by some operations.
	//
	// example:
	//
	// 99.90
	Rate *float64 `json:"rate,omitempty" xml:"rate,omitempty"`
	// The moderation scenario.
	//
	// example:
	//
	// antispam
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// The recommended subsequent operation. Valid values:
	//
	// 	- pass: The content passes the moderation.
	//
	// 	- review: The content needs to be manually reviewed again.
	//
	// 	- block: The content contains violations. We recommend that you delete or block the content.
	//
	// example:
	//
	// block
	Suggestion *string `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
}

func (s ImAuditResponseBodyTextResultsResultResults) String() string {
	return dara.Prettify(s)
}

func (s ImAuditResponseBodyTextResultsResultResults) GoString() string {
	return s.String()
}

func (s *ImAuditResponseBodyTextResultsResultResults) GetDetails() []*ImAuditResponseBodyTextResultsResultResultsDetails {
	return s.Details
}

func (s *ImAuditResponseBodyTextResultsResultResults) GetLabel() *string {
	return s.Label
}

func (s *ImAuditResponseBodyTextResultsResultResults) GetRate() *float64 {
	return s.Rate
}

func (s *ImAuditResponseBodyTextResultsResultResults) GetScene() *string {
	return s.Scene
}

func (s *ImAuditResponseBodyTextResultsResultResults) GetSuggestion() *string {
	return s.Suggestion
}

func (s *ImAuditResponseBodyTextResultsResultResults) SetDetails(v []*ImAuditResponseBodyTextResultsResultResultsDetails) *ImAuditResponseBodyTextResultsResultResults {
	s.Details = v
	return s
}

func (s *ImAuditResponseBodyTextResultsResultResults) SetLabel(v string) *ImAuditResponseBodyTextResultsResultResults {
	s.Label = &v
	return s
}

func (s *ImAuditResponseBodyTextResultsResultResults) SetRate(v float64) *ImAuditResponseBodyTextResultsResultResults {
	s.Rate = &v
	return s
}

func (s *ImAuditResponseBodyTextResultsResultResults) SetScene(v string) *ImAuditResponseBodyTextResultsResultResults {
	s.Scene = &v
	return s
}

func (s *ImAuditResponseBodyTextResultsResultResults) SetSuggestion(v string) *ImAuditResponseBodyTextResultsResultResults {
	s.Suggestion = &v
	return s
}

func (s *ImAuditResponseBodyTextResultsResultResults) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImAuditResponseBodyTextResultsResultResultsDetails struct {
	// The category of the risky content that the moderated text hits. Valid values:
	//
	// 	- spam: spam
	//
	// 	- ad: ad
	//
	// 	- politics: political content
	//
	// 	- terrorism: terrorist content
	//
	// 	- abuse: abuse
	//
	// 	- porn: pornographic content
	//
	// 	- flood: excessive junk content
	//
	// 	- contraband: prohibited content
	//
	// 	- meaningless: meaningless content
	//
	// 	- customized: custom content, such as a custom keyword
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The context information of the risky content that the moderated text hits.
	Contexts []*ImAuditResponseBodyTextResultsResultResultsDetailsContexts `json:"contexts,omitempty" xml:"contexts,omitempty" type:"Repeated"`
}

func (s ImAuditResponseBodyTextResultsResultResultsDetails) String() string {
	return dara.Prettify(s)
}

func (s ImAuditResponseBodyTextResultsResultResultsDetails) GoString() string {
	return s.String()
}

func (s *ImAuditResponseBodyTextResultsResultResultsDetails) GetLabel() *string {
	return s.Label
}

func (s *ImAuditResponseBodyTextResultsResultResultsDetails) GetContexts() []*ImAuditResponseBodyTextResultsResultResultsDetailsContexts {
	return s.Contexts
}

func (s *ImAuditResponseBodyTextResultsResultResultsDetails) SetLabel(v string) *ImAuditResponseBodyTextResultsResultResultsDetails {
	s.Label = &v
	return s
}

func (s *ImAuditResponseBodyTextResultsResultResultsDetails) SetContexts(v []*ImAuditResponseBodyTextResultsResultResultsDetailsContexts) *ImAuditResponseBodyTextResultsResultResultsDetails {
	s.Contexts = v
	return s
}

func (s *ImAuditResponseBodyTextResultsResultResultsDetails) Validate() error {
	if s.Contexts != nil {
		for _, item := range s.Contexts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImAuditResponseBodyTextResultsResultResultsDetailsContexts struct {
	// The term that the moderated text hits. If the text hits a term, the term is returned. If the text hits the algorithmic model, this parameter is not returned.
	//
	// example:
	//
	// Door-to-door service
	Context *string `json:"context,omitempty" xml:"context,omitempty"`
	// The code of the custom text library. This parameter is returned if the moderated text hits a term in the custom text library.
	//
	// example:
	//
	// 123456
	LibCode *string `json:"libCode,omitempty" xml:"libCode,omitempty"`
	// The name of the custom text library. This parameter is returned if the moderated text hits a term in the custom text library.
	//
	// example:
	//
	// Name of your custom text library
	LibName *string `json:"libName,omitempty" xml:"libName,omitempty"`
	// The position of the term that the moderated text hits in the original text.
	Positions []*string `json:"positions,omitempty" xml:"positions,omitempty" type:"Repeated"`
	// The behavior rule. This parameter is returned if the moderated text hits the behavior rule. Valid values:
	//
	// 	- user_id
	//
	// 	- ip
	//
	// 	- umid
	//
	// 	- content
	//
	// 	- similar_content
	//
	// 	- imei
	//
	// 	- imsi
	//
	// example:
	//
	// ip
	RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
}

func (s ImAuditResponseBodyTextResultsResultResultsDetailsContexts) String() string {
	return dara.Prettify(s)
}

func (s ImAuditResponseBodyTextResultsResultResultsDetailsContexts) GoString() string {
	return s.String()
}

func (s *ImAuditResponseBodyTextResultsResultResultsDetailsContexts) GetContext() *string {
	return s.Context
}

func (s *ImAuditResponseBodyTextResultsResultResultsDetailsContexts) GetLibCode() *string {
	return s.LibCode
}

func (s *ImAuditResponseBodyTextResultsResultResultsDetailsContexts) GetLibName() *string {
	return s.LibName
}

func (s *ImAuditResponseBodyTextResultsResultResultsDetailsContexts) GetPositions() []*string {
	return s.Positions
}

func (s *ImAuditResponseBodyTextResultsResultResultsDetailsContexts) GetRuleType() *string {
	return s.RuleType
}

func (s *ImAuditResponseBodyTextResultsResultResultsDetailsContexts) SetContext(v string) *ImAuditResponseBodyTextResultsResultResultsDetailsContexts {
	s.Context = &v
	return s
}

func (s *ImAuditResponseBodyTextResultsResultResultsDetailsContexts) SetLibCode(v string) *ImAuditResponseBodyTextResultsResultResultsDetailsContexts {
	s.LibCode = &v
	return s
}

func (s *ImAuditResponseBodyTextResultsResultResultsDetailsContexts) SetLibName(v string) *ImAuditResponseBodyTextResultsResultResultsDetailsContexts {
	s.LibName = &v
	return s
}

func (s *ImAuditResponseBodyTextResultsResultResultsDetailsContexts) SetPositions(v []*string) *ImAuditResponseBodyTextResultsResultResultsDetailsContexts {
	s.Positions = v
	return s
}

func (s *ImAuditResponseBodyTextResultsResultResultsDetailsContexts) SetRuleType(v string) *ImAuditResponseBodyTextResultsResultResultsDetailsContexts {
	s.RuleType = &v
	return s
}

func (s *ImAuditResponseBodyTextResultsResultResultsDetailsContexts) Validate() error {
	return dara.Validate(s)
}
